package hw10programoptimization

import (
	"bufio"
	"errors"
	"io"
	"strings"

	"github.com/valyala/fastjson"
)

type User struct {
	ID       int
	Name     string
	Username string
	Email    string
	Phone    string
	Password string
	Address  string
}

type DomainStat map[string]int

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	if domain == "" {
		return nil, errors.New("no domain provided")
	}
	DomainStat := make(DomainStat)
	var p fastjson.Parser
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		v, err := p.Parse(string(scanner.Bytes()))
		if err != nil {
			return nil, err
		}
		email := string(v.GetStringBytes("Email"))
		if email != "" && strings.HasSuffix(email, domain) {
			DomainStat[strings.ToLower(strings.SplitN(email, "@", 2)[1])]++
		}
	}
	return DomainStat, nil
}
