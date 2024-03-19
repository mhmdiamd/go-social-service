package communitymember

type handler struct {
  svc service
}

func newHandler(svc service) handler {
  return handler {
    svc : svc,
  }
}

