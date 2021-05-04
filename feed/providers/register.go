package providers

type ProvideFunc func(*Proxy) Provider

type Register struct {
	Domains []string
	Provide ProvideFunc
	Name    string
}

type RegisterMap map[string]*Register

func (r RegisterMap) AddRegister(name string, provide ProvideFunc, domains ...string) {
	r[name] = &Register{
		Name:    name,
		Domains: domains,
		Provide: provide,
	}
}

func (r RegisterMap) GetRegisterByURL(link string) *Register {
	reg := r[GENERIC]
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
