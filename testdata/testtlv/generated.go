// Code generated by radius-dict-gen. DO NOT EDIT.

package testtlv

import (
	"layeh.com/radius"
	"layeh.com/radius/rfc2865"
)

const (
	_TestVendor_VendorID = 9999
)

func _TestVendor_AddVendor(p *radius.Packet, typ byte, attr radius.Attribute) (err error) {
	var vsa radius.Attribute
	vendor := make(radius.Attribute, 2+len(attr))
	vendor[0] = typ
	vendor[1] = byte(len(vendor))
	copy(vendor[2:], attr)
	vsa, err = radius.NewVendorSpecific(_TestVendor_VendorID, vendor)
	if err != nil {
		return
	}
	p.Add(rfc2865.VendorSpecific_Type, vsa)
	return nil
}

func _TestVendor_GetsVendor(p *radius.Packet, typ byte) (values []radius.Attribute) {
	for _, attr := range p.Attributes[rfc2865.VendorSpecific_Type] {
		vendorID, vsa, err := radius.VendorSpecific(attr)
		if err != nil || vendorID != _TestVendor_VendorID {
			continue
		}
		for len(vsa) >= 3 {
			vsaTyp, vsaLen := vsa[0], vsa[1]
			if int(vsaLen) > len(vsa) || vsaLen < 3 {
				break
			}
			if vsaTyp == typ {
				values = append(values, vsa[2:int(vsaLen)])
			}
			vsa = vsa[int(vsaLen):]
		}
	}
	return
}

func _TestVendor_LookupVendor(p *radius.Packet, typ byte) (attr radius.Attribute, ok bool) {
	for _, a := range p.Attributes[rfc2865.VendorSpecific_Type] {
		vendorID, vsa, err := radius.VendorSpecific(a)
		if err != nil || vendorID != _TestVendor_VendorID {
			continue
		}
		for len(vsa) >= 3 {
			vsaTyp, vsaLen := vsa[0], vsa[1]
			if int(vsaLen) > len(vsa) || vsaLen < 3 {
				break
			}
			if vsaTyp == typ {
				return vsa[2:int(vsaLen)], true
			}
			vsa = vsa[int(vsaLen):]
		}
	}
	return nil, false
}

func _TestVendor_SetVendor(p *radius.Packet, typ byte, attr radius.Attribute) (err error) {
	for i := 0; i < len(p.Attributes[rfc2865.VendorSpecific_Type]); {
		vendorID, vsa, err := radius.VendorSpecific(p.Attributes[rfc2865.VendorSpecific_Type][i])
		if err != nil || vendorID != _TestVendor_VendorID {
			i++
			continue
		}
		for j := 0; len(vsa[j:]) >= 3; {
			vsaTyp, vsaLen := vsa[0], vsa[1]
			if int(vsaLen) > len(vsa[j:]) || vsaLen < 3 {
				i++
				break
			}
			if vsaTyp == typ {
				vsa = append(vsa[:j], vsa[j+int(vsaLen):]...)
			}
			j += int(vsaLen)
		}
		if len(vsa) > 0 {
			copy(p.Attributes[rfc2865.VendorSpecific_Type][i][4:], vsa)
			i++
		} else {
			p.Attributes[rfc2865.VendorSpecific_Type] = append(p.Attributes[rfc2865.VendorSpecific_Type][:i], p.Attributes[rfc2865.VendorSpecific_Type][i+i:]...)
		}
	}
	return _TestVendor_AddVendor(p, typ, attr)
}

const (
	TestTlv_Field1_Type radius.Type = 1
	TestTlv_Field2_Type radius.Type = 2
)

func TestTlv_Add(p *radius.Packet, attrs radius.Attributes) error {
	size := attrs.wireSize()
	b := make([]byte, size)
	attrs.encodeTo(b)
	return _TestVendor_AddVendor(p, 4, b)
}

func TestTlv_Get(p *radius.Packet) radius.Attributes {
	value, _ := TestTlv_Lookup(p)
	return value
}

func TestTlv_Gets(p *radius.Packet) (attrs_arr []radius.Attributes, err error) {
	for _, attr := range _TestVendor_GetsVendor(p, 4) {
		attr, err = ParseAttributes(attr)
		if err != nil {
			return
		}
		attrs_arr = append(attrs_arr, attr)
	}
	return
}

func TestTlv_Lookup(p *radius.Packet) (attrs radius.Attributes, err error) {
	a, ok := _TestVendor_LookupVendor(p, 4)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	attrs, err = ParseAttributes(a)
	return
}

func TestTlv_Set(p *radius.Packet, values []TestTlv) error {
	size := attrs.wireSize()
	b := make([]byte, size)
	attrs.encodeTo(b)
	return _TestVendor_AddVendor(p, 4, b)
}

func Field1_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _TestVendor_AddVendor(p, 1, a)
}

func Field1_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _TestVendor_AddVendor(p, 1, a)
}

func Field1_Get(p *radius.Packet) (value []byte) {
	value, _ = Field1_Lookup(p)
	return
}

func Field1_GetString(p *radius.Packet) (value string) {
	return string(Field1_Get(p))
}

func Field1_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _TestVendor_GetsVendor(p, 1) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func Field1_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _TestVendor_GetsVendor(p, 1) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func Field1_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _TestVendor_LookupVendor(p, 1)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func Field1_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _TestVendor_LookupVendor(p, 1)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func Field1_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _TestVendor_SetVendor(p, 1, a)
}

func Field1_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _TestVendor_SetVendor(p, 1, a)
}

type Field2 uint64

var Field2_Strings = map[Field2]string{}

func (a Field2) String() string {
	if str, ok := Field2_Strings[a]; ok {
		return str
	}
	return "Field2(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func Field2_Add(p *radius.Packet, value Field2) (err error) {
	a := radius.NewInteger64(uint64(value))
	return _TestVendor_AddVendor(p, 2, a)
}

func Field2_Get(p *radius.Packet) (value Field2) {
	value, _ = Field2_Lookup(p)
	return
}

func Field2_Gets(p *radius.Packet) (values []Field2, err error) {
	var i uint64
	for _, attr := range _TestVendor_GetsVendor(p, 2) {
		i, err = radius.Integer64(attr)
		if err != nil {
			return
		}
		values = append(values, Field2(i))
	}
	return
}

func Field2_Lookup(p *radius.Packet) (value Field2, err error) {
	a, ok := _TestVendor_LookupVendor(p, 2)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint64
	i, err = radius.Integer64(a)
	if err != nil {
		return
	}
	value = Field2(i)
	return
}

func Field2_Set(p *radius.Packet, value Field2) (err error) {
	a := radius.NewInteger64(uint64(value))
	return _TestVendor_SetVendor(p, 2, a)
}
