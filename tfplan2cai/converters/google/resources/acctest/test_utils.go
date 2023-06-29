package acctest

import (
	"errors"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func CheckDataSourceStateMatchesResourceState(dataSourceName, resourceName string) func(*terraform.State) error {
	return CheckDataSourceStateMatchesResourceStateWithIgnores(dataSourceName, resourceName, map[string]struct{}{})
}

func CheckDataSourceStateMatchesResourceStateWithIgnores(dataSourceName, resourceName string, ignoreFields map[string]struct{}) func(*terraform.State) error {
	return func(s *terraform.State) error {
		ds, ok := s.RootModule().Resources[dataSourceName]
		if !ok {
			return fmt.Errorf("can't find %s in state", dataSourceName)
		}

		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("can't find %s in state", resourceName)
		}

		dsAttr := ds.Primary.Attributes
		rsAttr := rs.Primary.Attributes

		errMsg := ""
		// Data sources are often derived from resources, so iterate over the resource fields to
		// make sure all fields are accounted for in the data source.
		// If a field exists in the data source but not in the resource, its expected value should
		// be checked separately.
		for k := range rsAttr {
			if _, ok := ignoreFields[k]; ok {
				continue
			}
			if k == "%" {
				continue
			}
			if dsAttr[k] != rsAttr[k] {
				// ignore data sources where an empty list is being compared against a null list.
				if k[len(k)-1:] == "#" && (dsAttr[k] == "" || dsAttr[k] == "0") && (rsAttr[k] == "" || rsAttr[k] == "0") {
					continue
				}
				errMsg += fmt.Sprintf("%s is %s; want %s\n", k, dsAttr[k], rsAttr[k])
			}
		}

		if errMsg != "" {
			return errors.New(errMsg)
		}

		return nil
	}
}

// This is a Printf sibling (Nprintf; Named Printf), which handles strings like
// Nprintf("Hello %{target}!", map[string]interface{}{"target":"world"}) == "Hello world!".
// This is particularly useful for generated tests, where we don't want to use Printf,
// since that would require us to generate a very particular ordering of arguments.
func Nprintf(format string, params map[string]interface{}) string {
	for key, val := range params {
		format = strings.Replace(format, "%{"+key+"}", fmt.Sprintf("%v", val), -1)
	}
	return format
}
