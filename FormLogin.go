// 由res2go自动生成，不要编辑。
package main

import (
    "github.com/ying32/govcl/vcl"
)

type TFormLogin struct {
    *vcl.TForm
    Label1     *vcl.TLabel
    Label2     *vcl.TLabel
    LabelEdit1 *vcl.TLabeledEdit
    LabelEdit2 *vcl.TLabeledEdit
    Button1    *vcl.TButton
    Button2    *vcl.TButton

    //::private::
    TFormLoginFields
}

var FormLogin *TFormLogin




// 以字节形式加载
// vcl.Application.CreateForm(formLoginBytes, &FormLogin)

func NewFormLogin(owner vcl.IComponent) (root *TFormLogin)  {
    vcl.CreateResForm(owner, formLoginBytes, &root)
    return
}

var formLoginBytes = []byte("\x54\x50\x46\x30\x0A\x54\x46\x6F\x72\x6D\x4C\x6F\x67\x69\x6E\x09\x46\x6F\x72\x6D\x4C\x6F\x67\x69\x6E\x04\x4C\x65\x66\x74\x02\x00\x03\x54\x6F\x70\x02\x00\x0B\x42\x6F\x72\x64\x65\x72\x49\x63\x6F\x6E\x73\x0B\x0C\x62\x69\x53\x79\x73\x74\x65\x6D\x4D\x65\x6E\x75\x0A\x62\x69\x4D\x69\x6E\x69\x6D\x69\x7A\x65\x00\x07\x43\x61\x70\x74\x69\x6F\x6E\x12\x02\x00\x00\x00\x7B\x76\x55\x5F\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x03\x23\x01\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\x8F\x01\x05\x43\x6F\x6C\x6F\x72\x07\x09\x63\x6C\x42\x74\x6E\x46\x61\x63\x65\x15\x43\x6F\x6E\x73\x74\x72\x61\x69\x6E\x74\x73\x2E\x4D\x69\x6E\x48\x65\x69\x67\x68\x74\x03\x4F\x01\x14\x43\x6F\x6E\x73\x74\x72\x61\x69\x6E\x74\x73\x2E\x4D\x69\x6E\x57\x69\x64\x74\x68\x03\xA1\x01\x0C\x46\x6F\x6E\x74\x2E\x43\x68\x61\x72\x73\x65\x74\x07\x0C\x41\x4E\x53\x49\x5F\x43\x48\x41\x52\x53\x45\x54\x0A\x46\x6F\x6E\x74\x2E\x43\x6F\x6C\x6F\x72\x07\x0C\x63\x6C\x57\x69\x6E\x64\x6F\x77\x54\x65\x78\x74\x0B\x46\x6F\x6E\x74\x2E\x48\x65\x69\x67\x68\x74\x02\xEE\x09\x46\x6F\x6E\x74\x2E\x4E\x61\x6D\x65\x12\x04\x00\x00\x00\xAE\x5F\x6F\x8F\xC5\x96\xD1\x9E\x0A\x46\x6F\x6E\x74\x2E\x53\x74\x79\x6C\x65\x0B\x00\x0E\x4F\x6C\x64\x43\x72\x65\x61\x74\x65\x4F\x72\x64\x65\x72\x08\x08\x50\x6F\x73\x69\x74\x69\x6F\x6E\x07\x0F\x70\x6F\x44\x65\x73\x6B\x74\x6F\x70\x43\x65\x6E\x74\x65\x72\x08\x4F\x6E\x43\x72\x65\x61\x74\x65\x07\x0A\x46\x6F\x72\x6D\x43\x72\x65\x61\x74\x65\x06\x4F\x6E\x53\x68\x6F\x77\x07\x08\x46\x6F\x72\x6D\x53\x68\x6F\x77\x0D\x50\x69\x78\x65\x6C\x73\x50\x65\x72\x49\x6E\x63\x68\x02\x78\x0A\x54\x65\x78\x74\x48\x65\x69\x67\x68\x74\x02\x18\x00\x06\x54\x4C\x61\x62\x65\x6C\x06\x4C\x61\x62\x65\x6C\x31\x04\x4C\x65\x66\x74\x03\x80\x00\x03\x54\x6F\x70\x02\x10\x05\x57\x69\x64\x74\x68\x03\xA2\x00\x06\x48\x65\x69\x67\x68\x74\x02\x21\x07\x43\x61\x70\x74\x69\x6F\x6E\x12\x06\x00\x00\x00\x70\x65\x6E\x63\x0C\x54\x65\x6B\xC4\x7E\xF6\x4E\x0C\x46\x6F\x6E\x74\x2E\x43\x68\x61\x72\x73\x65\x74\x07\x0F\x44\x45\x46\x41\x55\x4C\x54\x5F\x43\x48\x41\x52\x53\x45\x54\x0A\x46\x6F\x6E\x74\x2E\x43\x6F\x6C\x6F\x72\x07\x0C\x63\x6C\x57\x69\x6E\x64\x6F\x77\x54\x65\x78\x74\x0B\x46\x6F\x6E\x74\x2E\x48\x65\x69\x67\x68\x74\x02\xE5\x09\x46\x6F\x6E\x74\x2E\x4E\x61\x6D\x65\x06\x06\x54\x61\x68\x6F\x6D\x61\x0A\x46\x6F\x6E\x74\x2E\x53\x74\x79\x6C\x65\x0B\x00\x0A\x50\x61\x72\x65\x6E\x74\x46\x6F\x6E\x74\x08\x00\x00\x06\x54\x4C\x61\x62\x65\x6C\x06\x4C\x61\x62\x65\x6C\x32\x04\x4C\x65\x66\x74\x03\x8F\x00\x03\x54\x6F\x70\x03\x03\x01\x05\x57\x69\x64\x74\x68\x02\x69\x06\x48\x65\x69\x67\x68\x74\x02\x18\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x0B\x76\x32\x30\x32\x30\x2E\x30\x33\x2E\x30\x39\x00\x00\x0C\x54\x4C\x61\x62\x65\x6C\x65\x64\x45\x64\x69\x74\x0A\x4C\x61\x62\x65\x6C\x45\x64\x69\x74\x31\x04\x4C\x65\x66\x74\x03\x80\x00\x03\x54\x6F\x70\x02\x46\x05\x57\x69\x64\x74\x68\x03\xB4\x00\x06\x48\x65\x69\x67\x68\x74\x02\x20\x0F\x45\x64\x69\x74\x4C\x61\x62\x65\x6C\x2E\x57\x69\x64\x74\x68\x02\x3A\x10\x45\x64\x69\x74\x4C\x61\x62\x65\x6C\x2E\x48\x65\x69\x67\x68\x74\x02\x18\x11\x45\x64\x69\x74\x4C\x61\x62\x65\x6C\x2E\x43\x61\x70\x74\x69\x6F\x6E\x12\x04\x00\x00\x00\x28\x75\x37\x62\x0D\x54\x3A\x00\x0D\x4C\x61\x62\x65\x6C\x50\x6F\x73\x69\x74\x69\x6F\x6E\x07\x06\x6C\x70\x4C\x65\x66\x74\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x00\x00\x0C\x54\x4C\x61\x62\x65\x6C\x65\x64\x45\x64\x69\x74\x0A\x4C\x61\x62\x65\x6C\x45\x64\x69\x74\x32\x04\x4C\x65\x66\x74\x03\x80\x00\x03\x54\x6F\x70\x03\x80\x00\x05\x57\x69\x64\x74\x68\x03\xB4\x00\x06\x48\x65\x69\x67\x68\x74\x02\x20\x0F\x45\x64\x69\x74\x4C\x61\x62\x65\x6C\x2E\x57\x69\x64\x74\x68\x02\x3C\x10\x45\x64\x69\x74\x4C\x61\x62\x65\x6C\x2E\x48\x65\x69\x67\x68\x74\x02\x18\x11\x45\x64\x69\x74\x4C\x61\x62\x65\x6C\x2E\x43\x61\x70\x74\x69\x6F\x6E\x12\x07\x00\x00\x00\xC6\x5B\x20\x00\x20\x00\x20\x00\x20\x00\x01\x78\x3A\x00\x0C\x46\x6F\x6E\x74\x2E\x43\x68\x61\x72\x73\x65\x74\x07\x0C\x41\x4E\x53\x49\x5F\x43\x48\x41\x52\x53\x45\x54\x0A\x46\x6F\x6E\x74\x2E\x43\x6F\x6C\x6F\x72\x07\x0C\x63\x6C\x57\x69\x6E\x64\x6F\x77\x54\x65\x78\x74\x0B\x46\x6F\x6E\x74\x2E\x48\x65\x69\x67\x68\x74\x02\xEE\x09\x46\x6F\x6E\x74\x2E\x4E\x61\x6D\x65\x12\x04\x00\x00\x00\xAE\x5F\x6F\x8F\xC5\x96\xD1\x9E\x0A\x46\x6F\x6E\x74\x2E\x53\x74\x79\x6C\x65\x0B\x00\x0D\x4C\x61\x62\x65\x6C\x50\x6F\x73\x69\x74\x69\x6F\x6E\x07\x06\x6C\x70\x4C\x65\x66\x74\x0A\x50\x61\x72\x65\x6E\x74\x46\x6F\x6E\x74\x08\x0C\x50\x61\x73\x73\x77\x6F\x72\x64\x43\x68\x61\x72\x06\x01\x2A\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x01\x00\x00\x07\x54\x42\x75\x74\x74\x6F\x6E\x07\x42\x75\x74\x74\x6F\x6E\x31\x04\x4C\x65\x66\x74\x02\x38\x03\x54\x6F\x70\x03\xC0\x00\x05\x57\x69\x64\x74\x68\x02\x69\x06\x48\x65\x69\x67\x68\x74\x02\x31\x07\x43\x61\x70\x74\x69\x6F\x6E\x12\x02\x00\x00\x00\x7B\x76\x55\x5F\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x02\x07\x4F\x6E\x43\x6C\x69\x63\x6B\x07\x0C\x42\x75\x74\x74\x6F\x6E\x31\x43\x6C\x69\x63\x6B\x00\x00\x07\x54\x42\x75\x74\x74\x6F\x6E\x07\x42\x75\x74\x74\x6F\x6E\x32\x04\x4C\x65\x66\x74\x03\xDB\x00\x03\x54\x6F\x70\x03\xC0\x00\x05\x57\x69\x64\x74\x68\x02\x71\x06\x48\x65\x69\x67\x68\x74\x02\x31\x07\x43\x61\x70\x74\x69\x6F\x6E\x12\x02\x00\x00\x00\x00\x90\xFA\x51\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x03\x07\x4F\x6E\x43\x6C\x69\x63\x6B\x07\x0C\x42\x75\x74\x74\x6F\x6E\x32\x43\x6C\x69\x63\x6B\x00\x00\x00")
