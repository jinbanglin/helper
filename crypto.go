package helper

import (
  "crypto/md5"
  "crypto/sha1"
  "encoding/base64"
  "encoding/hex"
  "fmt"
  "io"
  "crypto/hmac"
  "net/url"
  "crypto/aes"
  "crypto/cipher"
  "errors"
)

//commonly used secret function
func MD5BySalt(src, salt string) string {
  hash := md5.New()
  io.WriteString(hash, src)
  io.WriteString(hash, salt)
  return hex.EncodeToString(hash.Sum(nil))
}

func MD5(b []byte) string {
  hash := md5.New()
  hash.Write(b)
  return base64.StdEncoding.EncodeToString(hash.Sum(nil))
}

func SHA1BySalt(src, salt string) (string, error) {
  t := sha1.New()
  _, err := io.WriteString(t, src)
  if err != nil {
    return "", err
  }
  _, err = io.WriteString(t, salt)
  if err != nil {
    return "", err
  }
  return fmt.Sprintf("%x", t.Sum(nil)), nil
}

func HmacSha1(src, key string) (string) {
  mac := hmac.New(sha1.New, String2Byte(key))
  mac.Write(String2Byte(src))
  //query := base64.StdEncoding.EncodeToString(mac.Sum(nil))
  //query = url.QueryEscape(query)
  return fmt.Sprintf("%x", mac.Sum(nil))
  //return query
}

func Sha1(query string, pri_key string) string {
  key := []byte(pri_key)
  mac := hmac.New(sha1.New, key)
  mac.Write([]byte(query))
  query = base64.StdEncoding.EncodeToString(mac.Sum(nil))
  query = url.QueryEscape(query)
  return query
}

func SHA1(src string) (string) {
  t := sha1.New()
  t.Write(String2Byte(src))
  return fmt.Sprintf("%x", t.Sum(nil))
}

type B64Encoding struct{ b *base64.Encoding }

func B64NewEncoding(s string) *B64Encoding {
  return &B64Encoding{b: base64.NewEncoding(s)}
}

func (b *B64Encoding) B64Encode(s string) string {
  return b.b.EncodeToString([]byte(s))
}

func (b *B64Encoding) B64Decode(s string) string {
  result, err := b.b.DecodeString(s)
  if err != nil {
    fmt.Println(err)
    return ""
  }
  return string(result)
}

func CBCDecrypter(keys string, endata string, ivs string) ([]byte, error) {
  seskey, err := base64.StdEncoding.DecodeString(keys)
  if err != nil {
    return []byte{}, err
  }
  data, err := base64.StdEncoding.DecodeString(endata)
  if err != nil {
    return []byte{}, err
  }
  iv, err := base64.StdEncoding.DecodeString(ivs)
  if err != nil {
    return []byte{}, err
  }
  return deCrypter(seskey, data, iv)

}

var eofDataError = errors.New("length not enough")
var multiBlockError = errors.New("ciphertext is not a multiple of the block size")

func deCrypter(key []byte, encodeData []byte, iv []byte) ([]byte, error) {
  block, err := aes.NewCipher(key)
  if err != nil {
    return []byte{}, err
  }
  if len(encodeData) < aes.BlockSize {
    return []byte{}, eofDataError
  }
  // CBC mode always works in whole blocks.
  if len(encodeData)%aes.BlockSize != 0 {
    return []byte{}, multiBlockError
  }
  mode := cipher.NewCBCDecrypter(block, iv)
  var data = make([]byte, len(encodeData))
  mode.CryptBlocks(data, encodeData)
  for i, v := range data {
    if v < 32 {
      data = data[:i]
      break
    }
  }
  return data, nil
}
