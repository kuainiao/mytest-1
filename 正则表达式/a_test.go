package 正则表达式

import (
	"testing"
	"regexp"
	"fmt"
)

func TestA(t *testing.T) {
	a()
}

func a() {
	aaa := []string{
		"popo.fhh.transport.m201809",
		"cpdt-tech-k8s-loda.picus-user-manager-prod.m201810",
		"cpdt-tech-k8s-loda.picus-user-manager-prod.m201811",
		"fhh-tech-k8s-loda.statistic-syncdata-prod.m201811",
		"cpdt-tech-k8s-loda.api-gateway-prod.m201811",
		"php11f-tech-k8s-loda.livesports-api-prod.m201811",
	}

	reg := regexp.MustCompile(`.(\w*)$`)
	for _, v := range aaa {
		a_a := reg.ReplaceAllString(v, ".*")
		fmt.Println(a_a)
	}

}
