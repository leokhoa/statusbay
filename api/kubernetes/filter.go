package kubernetes

import (
	"net/http"
	"statusbay/api/httpparameters"
	"strconv"
	"strings"
)

type FilterApplications struct {
	Offset        int
	Limit         int
	Clusters      []string
	Namespaces    []string
	Application   string
	Statuses      []string
	SortBy        string
	SortDirection string
	From          int
	To            int
	Distinct      bool
}

// FilterApplication Filter application by specific filters
func FilterApplication(req *http.Request) FilterApplications {

	offset, _ := strconv.Atoi(httpparameters.QueryParamWithDefault(req, "offset", "0"))
	limit, _ := strconv.Atoi(httpparameters.QueryParamWithDefault(req, "limit", "20"))
	cluster := httpparameters.QueryParamWithDefault(req, "cluster", "")
	namespace := httpparameters.QueryParamWithDefault(req, "namespace", "")
	status := httpparameters.QueryParamWithDefault(req, "status", "")
	sortBy := httpparameters.QueryParamWithDefault(req, "sortby", "time")
	sortDirection := httpparameters.QueryParamWithDefault(req, "sortdirection", "desc")
	from, _ := strconv.Atoi(httpparameters.QueryParamWithDefault(req, "from", "0"))
	to, _ := strconv.Atoi(httpparameters.QueryParamWithDefault(req, "to", "0"))
	distinct, _ := strconv.ParseBool(httpparameters.QueryParamWithDefault(req, "distinct", "false"))

	return FilterApplications{
		Limit:         limit,
		Offset:        offset,
		Clusters:      strings.Split(cluster, ","),
		Namespaces:    strings.Split(namespace, ","),
		Statuses:      strings.Split(status, ","),
		SortBy:        sortBy,
		SortDirection: sortDirection,
		From:          from,
		To:            to,
		Distinct:      distinct,
	}
}