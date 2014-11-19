package xapian

import (
  "fmt"
  "testing"
  "strings"
)

func TestVersion(t *testing.T) {
  if v := fmt.Sprintf("%d.%d.%d", Major_version(), Minor_version(), Revision()); Version_string() != v {
    t.Errorf("Unexpected version output");
  }
  if len(strings.Split(Version_string(),".")) != 3 {
    t.Errorf("Version_string() not X.Y.Z");
  }
  if strings.Split(Version_string(),".")[0] != "1" {
    t.Errorf("Version_string() not 1.Y.Z");
  } 
}

func TestGeneric1(t *testing.T) {
  stem := NewStem("english");
  if stem.Get_description() != "Xapian::Stem(english)" {
    t.Errorf("Unexpected str(stem)");
  }
  doc := NewDocument();
  doc.Set_data("a\x00b");
  if doc.Get_data() == "a" {
    t.Errorf("get_data+set_data truncates at a zero byte");
  }
  if doc.Get_data() != "a\x00b" {
    t.Errorf("get_data+set_data doesn't transparently handle a zero byte");
  }
  doc.Set_data("is there anybody out there?")
  doc.Add_term("XYzzy")
  doc.Add_posting(stem("is"), 1)
  doc.Add_posting(stem("there"), 2)
  doc.Add_posting(stem("anybody"), 3)
  doc.Add_posting(stem("out"), 4)
  doc.Add_posting(stem("there"), 5)
  //TODO: above fails, must implement operator() for Stem
}
