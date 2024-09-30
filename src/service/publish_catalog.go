package service

func (cs *catalogServiceInterface) PublishCatalog() {
	cs.repo.FindCatalog("1")
}
