package helmservice

import (
	helmrepo "githup.com/dierbei/go-helm-api/internal/repository/helm"
	"helm.sh/helm/v3/pkg/release"
)

var (
	myClusterApiServer = "https://172.16.0.220:8443"
	myClusterToken     = "ZXlKaGJHY2lPaUpTVXpJMU5pSXNJbXRwWkNJNkltaHdZbXhOZURka1dYVmxUREp1VjB0MlVGQkthemw0WXpBd05YTkxVekpMUkRWRlNHaHNWRzFvVXpnaWZRLmV5SnBjM01pT2lKcmRXSmxjbTVsZEdWekwzTmxjblpwWTJWaFkyTnZkVzUwSWl3aWEzVmlaWEp1WlhSbGN5NXBieTl6WlhKMmFXTmxZV05qYjNWdWRDOXVZVzFsYzNCaFkyVWlPaUprWldaaGRXeDBJaXdpYTNWaVpYSnVaWFJsY3k1cGJ5OXpaWEoyYVdObFlXTmpiM1Z1ZEM5elpXTnlaWFF1Ym1GdFpTSTZJblJsYzNSamJIVnpkR1Z5YzJFdGRHOXJaVzR0ZUhKdVl6SWlMQ0pyZFdKbGNtNWxkR1Z6TG1sdkwzTmxjblpwWTJWaFkyTnZkVzUwTDNObGNuWnBZMlV0WVdOamIzVnVkQzV1WVcxbElqb2lkR1Z6ZEdOc2RYTjBaWEp6WVNJc0ltdDFZbVZ5Ym1WMFpYTXVhVzh2YzJWeWRtbGpaV0ZqWTI5MWJuUXZjMlZ5ZG1salpTMWhZMk52ZFc1MExuVnBaQ0k2SWpWbFpESXpZelV4TFRFMU1XSXROREU0WkMxaU9EazRMVFZqWW1Nd01UWmlNREF4WVNJc0luTjFZaUk2SW5ONWMzUmxiVHB6WlhKMmFXTmxZV05qYjNWdWREcGtaV1poZFd4ME9uUmxjM1JqYkhWemRHVnljMkVpZlEuQm04MjdEWGtmUG9VdnFNNHd2MGVsdE5hMHE2RlZLTUMxTzl4ODFoTkwzbE9oVFNkUEoyZDl1Yjh3VWtRa0tzYXpLSVc1OG1ybGpOTGJoNUsyeDVzTFNkVnBJLUpsTU9oekpKZTZJM2c0bXEwdExPblEyRjJNcFZvTzRUdWhMUTFmWE5zWHo4bGhUSW01ekRaeW4wMnA1emQ0a3FIQ1BHcFVWeDRvWEtHT1BHbkN3WnJTTTdMdVc5R1Q2RWc2RmJMVl93bEpVak9ha1QxNW10dmxvQmJuNkNaU0V6ODVpT0UxQ3V1cW42VTdpRV9XZkdWVXJXbjQ3eE1VSUtsd2RiY1dCczE1N0JpWkNNMnhpNmNXYzdnRGxTYi1VNjFQUjlrSUMyN0EtNkZEZXZ3ZHJsRGJQdnlvZmh3VE42ZHRQbWRwTV81MnE5d0twU3pKQmhlbUNGaER3"
	myClusterCa        = "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUN5RENDQWJDZ0F3SUJBZ0lCQURBTkJna3Foa2lHOXcwQkFRc0ZBREFWTVJNd0VRWURWUVFERXdwcmRXSmwKY201bGRHVnpNQjRYRFRJd01EVXdOakV6TXpVd05Wb1hEVE13TURVd05ERXpNelV3TlZvd0ZURVRNQkVHQTFVRQpBeE1LYTNWaVpYSnVaWFJsY3pDQ0FTSXdEUVlKS29aSWh2Y05BUUVCQlFBRGdnRVBBRENDQVFvQ2dnRUJBTTBECm54bHlhOUFQNis2UllzQ1FtNkl4YzhubVR3Wm9NV1dkeUUvaVFPblZuY3Z5VE1DeDRJWVBURUNxY3g3V0JiLy8Kbk4wYjlxUXFxbHcxTm9jd0pSdTRvVEFJSitCcDljU013ZTQycTR2c3JrbTdHYWJlb0hRUFFtZDl2cGNVVXl5aApDZVZHRmhXOEFXbmR3RURlNVNhYjdjQ2dUUXI5MXRzK2dnUTVreDRjMkFXSkpFOWFtTm01YnpCeVJoWUN5cDdnCkpRb2VMUXltSFVMRENLZDk1RHA2a0tQTWhRVkdqbWRDQkM5OTJGaFhrcGlqU1FDbUFQdGpEeU5velhaMFU2dzkKMGZVank0RG8yUWZhTHVZN0UzcDQvREh5L2Q5TFF1Zy95NVBjUlRGVGVuOXNhNjZwREFac3hBUDFpOFlRQUhNRQo5U0pjaTZwbUV6L0JCK20zUkZVQ0F3RUFBYU1qTUNFd0RnWURWUjBQQVFIL0JBUURBZ0trTUE4R0ExVWRFd0VCCi93UUZNQU1CQWY4d0RRWUpLb1pJaHZjTkFRRUxCUUFEZ2dFQkFEZkdWTTVIQzJDaTVhQkdYV1RtMnVTSVhyeFkKdlJJTlBJaDJGV0JyMHBkWUtra2pzc1UveDZhRmkwNDJjZE84aWtiaGxWKzdzZ1ZBYncyNnowSW8rK1ZrOHhSUAppanB3YmltdkNLN0t2T0hMTVJ5d3ZHajYralUxc1EraHRMenkwUk41WFZOVkVkTVlUdUprbWR3eDBMY2Q3Tnd4CnpzMmxaYnVvZGhhMTFzWENTc05yZHUzY2k5U2JVTFltbEVNZHdwRU0wRzJBUkFkdzdFWDJhQm9KS21YbnZiSGMKa3IvYVJPSVNGZjFHNWYrNm95b1o1d2c4WDR6QWJwZUN6SmZqK1BJVktFMXR3RC9ZczVCc0lFOEg5MmIyTTdLSgplVE9PZERQSG11MURrdG5xTnhpRmprZ2xwWFlFdTV5aU1Rbm9zd00rN3lwa3BiSmhtb0g1UFJkTTVWMD0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="
)

//var myClusterToken = "ZXlKaGJHY2lPaUpTVXpJMU5pSXNJbXRwWkNJNklsVjZjemxUV25kT1lsQmpVMjFYTUdKQ1lVOW5OM0p6ZVZaeUxXSk1aMlpIVGpnNFJXZEhUelJSVTJzaWZRLmV5SnBjM01pT2lKcmRXSmxjbTVsZEdWekwzTmxjblpwWTJWaFkyTnZkVzUwSWl3aWEzVmlaWEp1WlhSbGN5NXBieTl6WlhKMmFXTmxZV05qYjNWdWRDOXVZVzFsYzNCaFkyVWlPaUowWld0MGIyNHRjR2x3Wld4cGJtVnpJaXdpYTNWaVpYSnVaWFJsY3k1cGJ5OXpaWEoyYVdObFlXTmpiM1Z1ZEM5elpXTnlaWFF1Ym1GdFpTSTZJakUxTVRBMU1ERTFOVGt4TFhObGNuWnBZMlZoWTJOdmRXNTBMWFJ2YTJWdUxYaHJOMjF5SWl3aWEzVmlaWEp1WlhSbGN5NXBieTl6WlhKMmFXTmxZV05qYjNWdWRDOXpaWEoyYVdObExXRmpZMjkxYm5RdWJtRnRaU0k2SWpFMU1UQTFNREUxTlRreExYTmxjblpwWTJWaFkyTnZkVzUwSWl3aWEzVmlaWEp1WlhSbGN5NXBieTl6WlhKMmFXTmxZV05qYjNWdWRDOXpaWEoyYVdObExXRmpZMjkxYm5RdWRXbGtJam9pTVRFd05Ua3lZVEV0WWpobE15MDBaR1kzTFRrMk5UTXRNV05qWW1aa1pEQTJNelE0SWl3aWMzVmlJam9pYzNsemRHVnRPbk5sY25acFkyVmhZMk52ZFc1ME9uUmxhM1J2Ymkxd2FYQmxiR2x1WlhNNk1UVXhNRFV3TVRVMU9URXRjMlZ5ZG1salpXRmpZMjkxYm5RaWZRLlRxdVBNUmVSaHZ2TmFoMU5yQzVvbEhlelVVampTNURWQ014U3dyTVVBcWxmZEQzajdHalJKNGdkbmJIWGtCdzFJNXpnOEtpQTJHTVZmZWR4RmFRUDZaU0VFckI4TVRjckJjbUplVkJ1M1FydjE4N1kzdGtmcFNyay16MVMtU2l6eXVpNWNaM2hmYXpUMHpCS09tQ3dMUTktMWV3WDRxNHJpbUVUaXFnVDd5aG5hZ2JBYlgySUJDaWV4YXlvUkdSNlJJdEszMFZLZUZidmd5bW5vLTRZSFpEYUhRUmtzSDY0Sk94dUNXWlFkWjdGU3RTczlmRU13T1pLODdxTEx5ZWJnZG1OZGo3MEFxUjhyWE1fOE8wSjl4TVF3TGxqNmVxdS00V2hkcGs1ckl2MTl3OGpqQzRuTDVzSEtPbVZoVDVGUVM5RHotWTVCckxKNG5fZTVGTHNBdw=="
//var myClusterCa = "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUMvakNDQWVhZ0F3SUJBZ0lCQURBTkJna3Foa2lHOXcwQkFRc0ZBREFWTVJNd0VRWURWUVFERXdwcmRXSmwKY201bGRHVnpNQjRYRFRJeU1ETXlOekEyTURFd05Gb1hEVE15TURNeU5EQTJNREV3TkZvd0ZURVRNQkVHQTFVRQpBeE1LYTNWaVpYSnVaWFJsY3pDQ0FTSXdEUVlKS29aSWh2Y05BUUVCQlFBRGdnRVBBRENDQVFvQ2dnRUJBTDZnCkNzR3Q4NUdyMm9KYUdZRys3d25VWDl1MHVoa0ZqYlpJZk1JR3loemdRZmVtRGhVL3pMV2ZDMFJaV0s1OHZtVEgKY1FHeDZvNGxvMC80SU91RTNSaXlSR1N4dUE1NkNpYy8rcVR6WnhtaERnSkcrSkMyRGs2QnZhZEkycDlGNDIyKwoza2VnZm1Pc3M3clRSaXZJZEdyYXp2S01nVm5WcFQxS2RiM1lEZ0loRkNlZjJiUjR2Y3ljU1VjZjRvaHc0S28xCmE2VHlCOWlBQ1Vacmo0UmdycThRYkR6VzBPWWdOVE45QURsVXBoRWJramNaeU9XYXMyWndEK2lEcXpnRHB6b0MKQzJ3bUN3dnhZWUlaZEE1dW9aaGxmRndmREJBeVVlV2JJZ200MUhRT0tZM0lOa3h0d3FXSEJEWEJsVVo3OWVvVwo3TWNTWDB0NnR4OUZpdElHY1djQ0F3RUFBYU5aTUZjd0RnWURWUjBQQVFIL0JBUURBZ0trTUE4R0ExVWRFd0VCCi93UUZNQU1CQWY4d0hRWURWUjBPQkJZRUZBamhUTUUvZjc1TXVnbUw0K0lKOGlEMzJtdWtNQlVHQTFVZEVRUU8KTUF5Q0NtdDFZbVZ5Ym1WMFpYTXdEUVlKS29aSWh2Y05BUUVMQlFBRGdnRUJBRFA5Q1BJTTdLVEYwTldINHZNUQowNlpENUNXa3Y3UmVaK0doalJ4K1M1bFZTd1pzeDhpV0N1RUNrRGVSWURield5cFVRVUVqMmFTTlFUY0tCaDkwCklaeDRhTk1HbE9scWxkdXBWVmRnVGxSOVBpSCtaaldENXluL3pMa1hSUGJJVXlNUktjOWlXcDNlM3hLRnRYYVMKOXBJUGRrNnM1NDF5cUhpWGRRYVZ1RHdUTndmcWtRcnczVW9pNlEyYkh6WER3VUpyZjFZSjVuUyt1dkg4ZHBTLwpDNkRiRGlqemZkaVlrc0pXTXJUSTVhYjREdVRvVjFwS1FOQmJJN1FMeTkzS1ZxSlQxYnlNVkdQQm1BSStYMHFICjZ2c2x0WVAxVlhibmFvQncyQ2VMa3JuNGYxbVU5T054NDNPYkhWOFVHbll3TXV5K2IvelBrSXcyUWU2TzVtQmcKbzcwPQotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg=="
//var myClusterApiServer = "https://175.24.198.168:6443"

var _ Service = (*service)(nil)

type Service interface {
	i()

	ListRepositoryChart(repoName, version, keyword, versions string) (*helmrepo.ListRepoChartResponse, error)

	AddRepository(req *helmrepo.AddRepositoryRequest) error

	UpdateRepository(req *helmrepo.UpdateRepositoryRequest) (*helmrepo.UpdateRepositoryResponse, error)

	SelectRepository(search *helmrepo.Repository) (*helmrepo.Repository, error)

	InstallOrUpgradeRelease(namespace, repository, release, chart, version string) (*release.Release, error)

	ListRepository(page, pageSize int) (*helmrepo.ListRepositoryResponse, error)

	UninstallRelease(namespace, release string) error

	GetChartInfo(chart, info, version string) (interface{}, error)

	InitRepos() error

	ShowReleaseInfo(release, namespace, info, output string) (interface{}, error)
}

type service struct {
}

func New() Service {
	return &service{}
}

func (s *service) i() {}
