package providers

type Register struct {
	Domains []string
	Provide func(*Proxy) Provider
	Name    string
}

type RegisterMap map[string]*Register

func (r RegisterMap) GetRegisterByURL(link string) *Register {
	reg := r["generic"]
	domain := getDomain(link)
	for _, v := range r {
		for _, d := range v.Domains {
			if d == domain {
				return v
			}
		}
	}
	return reg
}

func (r RegisterMap) GetRegisterByKey(key string) *Register {
	var reg *Register
	for k, v := range r {
		if k == key {
			return v
		}
	}
	return reg
}
