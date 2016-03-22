package guber

type ReplicationControllers struct {
	client    *Client
	Namespace string
}

func (r ReplicationControllers) DomainName() string {
	return ""
}

func (r ReplicationControllers) ApiGroup() string {
	return "api"
}

func (r ReplicationControllers) ApiVersion() string {
	return "v1"
}

func (r ReplicationControllers) ApiName() string {
	return "replicationcontrollers"
}

func (r ReplicationControllers) Kind() string {
	return "ReplicationController"
}

func (r *ReplicationControllers) Create(e *ReplicationController) (*ReplicationController, error) {
	if err := r.client.Post().Resource(r).Namespace(r.Namespace).Entity(e).Do().Into(e); err != nil {
		return nil, err
	}
	return e, nil
}

func (r *ReplicationControllers) List(q *QueryParams) (*ReplicationControllerList, error) {
	list := new(ReplicationControllerList)
	err := r.client.Get().Resource(r).Namespace(r.Namespace).Query(q).Do().Into(list)
	return list, err
}

func (r *ReplicationControllers) Get(name string) (*ReplicationController, error) {
	e := new(ReplicationController)
	if err := r.client.Get().Resource(r).Namespace(r.Namespace).Name(name).Do().Into(e); err != nil {
		return nil, err
	}
	return e, nil
}

func (r *ReplicationControllers) Update(name string, e *ReplicationController) (*ReplicationController, error) {
	if err := r.client.Patch().Resource(r).Namespace(r.Namespace).Name(name).Entity(e).Do().Into(e); err != nil {
		return nil, err
	}
	return e, nil
}

func (r *ReplicationControllers) Delete(name string) (found bool, err error) {
	req := r.client.Delete().Resource(r).Namespace(r.Namespace).Name(name).Do()
	return req.found, req.err
}
