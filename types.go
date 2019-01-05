package core

type Router interface{
  Serve(int)
}

type Headers map[string]string
