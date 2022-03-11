package main

type user struct {
	id       int
	username int
	password int
}

type userservice struct {
	t []user
	// penamaan variable t tidak mendesktipsikan apa yang variable ini wakilkan
}

/*
biasanya penamaan function yang lebih dari satu kata untuk kata seanjutnya diawali dengan huruf kapital
maka penulisan dari function dibawah dapat menjadi getAllUsers
*/
func (u userservice) getallusers() []user {
	return u.t
}

// karena kode ini seluruhnya tentang user, maka sebenarnya penamaan dibawah dapat menjadi lebih sederhana
// contohnya getId
func (u userservice) getuserbyid(id int) user {
	for _, r := range u.t {
		// variable "r" disini juga tidak mudah untuk dipahami
		if id == r.id {
			return r
		}
	}
	return user{}
}

// kode ini tidak mempunyai satupun komentar yang menjelaskan apa yang akan dilakukan
