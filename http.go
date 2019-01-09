package helper

import (
  "bytes"
  "crypto/tls"
  "io"
  "io/ioutil"
  "net"
  "net/http"
  "net/url"
  "strings"
  "time"
)

var client *http.Client

func HTTPInstance() *http.Client {
  if client == nil {
    client = &http.Client{
      Transport: &http.Transport{
        Dial:         printLocalDial,
        MaxIdleConns: 100,
      },
    }
  }
  return client
}

func HTTPTLSInstance(certFile, keyFile string) *http.Client {
  cert, err := tls.LoadX509KeyPair(certFile, keyFile)
  if err != nil {
    panic(err)
  }
  return &http.Client{
    Transport: &http.Transport{
      Proxy:                 http.ProxyFromEnvironment,
      Dial:                  printLocalDial,
      TLSHandshakeTimeout:   10 * time.Second,
      ExpectContinueTimeout: 1 * time.Second,
      TLSClientConfig: &tls.Config{
        Certificates: []tls.Certificate{cert},
      },
    },
    Timeout: 60 * time.Second,
  }
}

func printLocalDial(network, addr string) (net.Conn, error) {
  dial := &net.Dialer{
    Timeout:   5 * time.Second,
    KeepAlive: 30 * time.Second,
  }
  conn, err := dial.Dial(network, addr)
  if err != nil {
    return conn, err
  }
  //print what host to use
  //fmt.Println("connect done, use", conn.LocalAddr().String())
  return conn, err
}

func PostUrlencoded(address string, isKeep bool, data map[string]string) ([]byte, error) {
  u := url.Values{}
  for k, v := range data {
    u.Add(k, v)
  }
  resp, err := HTTPInstance().Post(address, "application/x-www-form-urlencoded", strings.NewReader(u.Encode()))
  if err != nil {
    return nil, err
  }
  if isKeep {
    defer func() {
      resp.Body.Close()
      io.Copy(ioutil.Discard, resp.Body)
    }()
  } else {
    resp.Close = true
  }
  return ioutil.ReadAll(resp.Body)
}

func PostRawJson(address string, isKeep bool, data []byte) ([]byte, error) {
  resp, err := HTTPInstance().Post(address, "application/json;utf-8", bytes.NewBuffer(data))
  if err != nil {
    return nil, err
  }
  if isKeep {
    defer func() {
      resp.Body.Close()
      io.Copy(ioutil.Discard, resp.Body)
    }()
  } else {
    resp.Close = true
  }
  return ioutil.ReadAll(resp.Body)
}

func Post(address, contentType string, isKeep bool, data []byte) ([]byte, error) {
  resp, err := HTTPInstance().Post(address, contentType, bytes.NewBuffer(data))
  if err != nil {
    return nil, err
  }
  if isKeep {
    defer func() {
      resp.Body.Close()
      io.Copy(ioutil.Discard, resp.Body)
    }()
  } else {
    resp.Close = true
  }
  return ioutil.ReadAll(resp.Body)
}

func GetUrlencoded(address string, isKeep bool, data map[string]string) ([]byte, error) {
  u, _ := url.Parse(address)
  q := u.Query()
  for k, v := range data {
    q.Set(k, v)
  }
  u.RawQuery = q.Encode()
  resp, err := HTTPInstance().Get(u.String())
  if err != nil {
    return nil, err
  }
  if isKeep {
    defer func() {
      resp.Body.Close()
      io.Copy(ioutil.Discard, resp.Body)
    }()
  } else {
    resp.Close = true
  }
  return ioutil.ReadAll(resp.Body)
}

func Get(address string, isKeep bool) ([]byte, error) {
  u, _ := url.Parse(address)

  resp, err := HTTPInstance().Get(u.String())
  if err != nil {
    return nil, err
  }
  if isKeep {
    defer func() {
      resp.Body.Close()
      io.Copy(ioutil.Discard, resp.Body)
    }()
  } else {
    resp.Close = true
  }
  return ioutil.ReadAll(resp.Body)
}
