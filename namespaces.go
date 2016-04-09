package guber

type Namespaces struct {
	client *Client
}

func (c *Namespaces) New() *Namespace {
	return &Namespace{
		collection: c,
	}
}

func (c Namespaces) DomainName() string {
	return ""
}

func (c Namespaces) APIGroup() string {
	return "api"
}

func (c Namespaces) APIVersion() string {
	return "v1"
}

func (c Namespaces) APIName() string {
	return "namespaces"
}

func (c Namespaces) Kind() string {
	return "Namespace"
}

func (c *Namespaces) Create(e *Namespace) (*Namespace, error) {
	r := c.New()
	if err := c.client.Post().Collection(c).Entity(e).Do().Into(r); err != nil {
		return nil, err
	}
	return r, nil
}

func (c *Namespaces) Query(q *QueryParams) (*NamespaceList, error) {
	list := new(NamespaceList)
	if err := c.client.Get().Collection(c).Query(q).Do().Into(list); err != nil {
		return nil, err
	}
	for _, r := range list.Items {
		r.collection = c
	}
	return list, nil
}

func (c *Namespaces) List() (*NamespaceList, error) {
	list := new(NamespaceList)
	if err := c.client.Get().Collection(c).Do().Into(list); err != nil {
		return nil, err
	}
	for _, r := range list.Items {
		r.collection = c
	}
	return list, nil
}

func (c *Namespaces) Get(name string) (*Namespace, error) {
	r := c.New()
	req := c.client.Get().Collection(c).Name(name).Do()
	if err := req.Into(r); err != nil {
		return nil, err
	}
	if req.found {
		return r, nil
	}
	return nil, nil
}

func (c *Namespaces) Update(name string, r *Namespace) (*Namespace, error) {
	if err := c.client.Patch().Collection(c).Name(name).Entity(r).Do().Into(r); err != nil {
		return nil, err
	}
	return r, nil
}

func (c *Namespaces) Delete(name string) (found bool, err error) {
	req := c.client.Delete().Collection(c).Name(name).Do()
	return req.found, req.err
}

// Resource-level

func (r *Namespace) Reload() (*Namespace, error) {
	return r.collection.Get(r.Metadata.Name)
}

func (r *Namespace) Save() error {
	_, err := r.collection.Update(r.Metadata.Name, r)
	return err
}

func (r *Namespace) Delete() error {
	_, err := r.collection.Delete(r.Metadata.Name)
	return err
}
