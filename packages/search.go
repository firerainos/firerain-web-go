package packages

func GetArch() []Option {
	return GetSelectText("#id_arch")
}

func GetRepository() []Option {
	return GetSelectText("#id_repo")
}

func GetMaintainer() []Option {
	return GetSelectText("#id_maintainer")
}

func GetFlagged() []Option {
	return GetSelectText("#id_flagged")
}