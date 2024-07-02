package tests
import "src/crypto_utils"
import "testing"


// This should return an error
func TestGetHashedPassword_Empty(t *testing.T){
	str := ""
	b64hash,err := crypto_utils.GetHashedPassword(str)
	if(err == nil){
		t.Fatalf("GetHashedPassword() Empty Password did not fail !: %v", err)
	}
	_ = b64hash
}


// This should not return an error
func TestGetHashedPassword(t *testing.T){
	str := "test1234"
	b64hash,err := crypto_utils.GetHashedPassword(str)
	_ = b64hash
	if(err != nil){
		t.Fatalf("GetHashedPassword() failed: %v", err)
	}
}

// This should not return an error
func TestGetHashedPassword_Large(t *testing.T){
	// Str is bigger than 64 which should throw an error
	str := "7cwsrd63QIAXGlZorJzNdIeBuB4ClWLC2NUotWGhtEn3fhE8WUEmEWrvavj4mj0qN"
	b64hash,err := crypto_utils.GetHashedPassword(str)
	if(err == nil){
		t.Fatalf("GetHashedPassword() with big password didnt fail!: %v", err)
	}
	_ = b64hash
}

func TestCheckPasswordCorrect_Correct(t *testing.T){
	str := "test1234"
	b64hash,err := crypto_utils.GetHashedPassword(str)
	if(err != nil){
		t.Fatalf("CheckPasswordCorrect() failed because of GetHashedPassword!: %v", err)
	}
	b,e := crypto_utils.CheckPasswordCorrect(str,b64hash)
	if(e != nil){
		t.Fatalf("CheckPasswordCorrect() failed because of CheckPasswordCorrect error!: %v", err)
	}

	if(b == false){
		t.Fatalf("CheckPasswordCorrect() failed because of CheckPasswordCorrect not correctly comparing the passwords!: %v", err)
	}
}


func TestCheckPasswordCorrect_Wrong(t *testing.T){
	str := "test1234"
	b64hash,err := crypto_utils.GetHashedPassword(str)
	if(err != nil){
		t.Fatalf("CheckPasswordCorrect() failed because of GetHashedPassword!: %v", err)
	}
	b,e := crypto_utils.CheckPasswordCorrect("test11111111",b64hash)
	if(e != nil){
		t.Fatalf("CheckPasswordCorrect() failed because of CheckPasswordCorrect error!: %v", err)
	}

	if(b != false){
		t.Fatalf("CheckPasswordCorrect() failed because of CheckPasswordCorrect not correctly comparing the passwords!: %v", err)
	}
}