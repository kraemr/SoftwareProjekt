package crypto_utils;
import (
	"golang.org/x/crypto/argon2"
	"encoding/base64"
	"fmt"
	"crypto/rand"
	"errors"
	"strings"
	"crypto/subtle"
)

var (
    ErrInvalidHash         = errors.New("the encoded hash is not in the correct format")
    ErrIncompatibleVersion = errors.New("incompatible version of argon2")
    ErrInvalidPassword         = errors.New("the Password is empty or has invalid characters")

)

type params struct{
	memory      uint32
    iterations  uint32
    parallelism uint8
    saltLength  uint32
    keyLength   uint32
}

const iterations = 1
const memory = 2048 // 2048 KB
const parallelism = 2
const keyLength  =32
const saltLength = 16

type hash struct {
    
}





func decodeHash(encodedHash string) (p *params, salt, hash []byte, err error) {
    vals := strings.Split(encodedHash, "$")
    if len(vals) != 6 {
        return nil, nil, nil, ErrInvalidHash
    }
    var version int
    _, err = fmt.Sscanf(vals[2], "v=%d", &version)
    if err != nil {
        return nil, nil, nil, err
    }
    if version != argon2.Version {
        return nil, nil, nil, ErrIncompatibleVersion
    }

    p = &params{}
    _, err = fmt.Sscanf(vals[3], "m=%d,t=%d,p=%d", &p.memory, &p.iterations, &p.parallelism)
    if err != nil {
        return nil, nil, nil, err
    }

    salt, err = base64.RawStdEncoding.Strict().DecodeString(vals[4])
    if err != nil {
        return nil, nil, nil, err
    }
    p.saltLength = uint32(len(salt))

    hash, err = base64.RawStdEncoding.Strict().DecodeString(vals[5])
    if err != nil {
        return nil, nil, nil, err
    }
    p.keyLength = uint32(len(hash))

    return p, salt, hash, nil
}

func generateSalt(n uint32) ([]byte, error) {
    b := make([]byte, n)
    _, err := rand.Read(b)
    if err != nil {
        return nil, err
    }
    return b, nil
}

func formatBase64Argon2(hash []byte,salt []byte) string{
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)
	encodedHash := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, memory, iterations, parallelism, b64Salt, b64Hash)
	return encodedHash
}

func getHashedPasswordWithParams(password string,argon2Params params) (string,error){
    if( len(password) == 0 || len(password) > 64){
        return "",ErrInvalidPassword
    }
    
    salt, err := generateSalt(saltLength)
    if err != nil {
        return "", err
    }
    hash := argon2.IDKey([]byte(password), salt, argon2Params.iterations, argon2Params.memory, argon2Params.parallelism, argon2Params.keyLength)
	b64Hash := formatBase64Argon2(hash,salt);
	return b64Hash, nil
}

// returns hashed Password in base64
// uses defaults defined here
func GetHashedPassword(password string) (string,error) {
    salt, err := generateSalt(saltLength)
    if err != nil {
        return "", err
    }
    hash := argon2.IDKey([]byte(password), salt, iterations, memory, parallelism, keyLength)
	b64Hash := formatBase64Argon2(hash,salt);
	return b64Hash, nil
}

// This receives a base64 encoded and formatted hashed password
// with salt and parameters included 
// This enables us to hash the cleartext password with the same parameters and salt
func CheckPasswordCorrect(password string, hashedPassword string) (bool,error){
	p, salt, hash, err := decodeHash(hashedPassword)
    if err != nil {
        return false, err
    }
	otherHash := argon2.IDKey([]byte(password), salt, p.iterations, p.memory, p.parallelism, p.keyLength)
	if subtle.ConstantTimeCompare(hash, otherHash) == 1 {
		return true, nil
    }
	return false, nil
}