package response

type GeneralResources struct {
	Origins    []Origin   `json:"origins"`
	Priorities []Priority `json:"priorities"`
	Types      []Type     `json:"types"`
}
