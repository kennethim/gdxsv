package main

import (
	"bytes"
	"testing"
)

func TestMessageBodyReader_ReadEncryptedString(t *testing.T) {
	type fields struct {
		seq uint16
		r   *bytes.Reader
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"example1",
			fields{401, bytes.NewReader([]byte{0x00, 0x22, 0x14, 0x80, 0x06, 0x37, 0x0b, 0x71, 0x1c, 0x4b, 0x0f, 0x3b, 0x1e, 0x53, 0x1e, 0x11, 0x04, 0x45, 0x17, 0x40, 0x16, 0x23, 0x1b, 0x54, 0x0c, 0x20, 0x1f, 0x4b, 0x2e, 0x5f, 0x23, 0x60, 0x34, 0x19, 0x27, 0x0e, 0x00})},
			"こんにちは初めましてごめんなさい",
		},
		{
			"example2",
			fields{402, bytes.NewReader([]byte{0x00, 0x26, 0x14, 0xb3, 0x06, 0xf5, 0x0b, 0xdf, 0x1c, 0x0e, 0x0d, 0xc0, 0x1e, 0xf5, 0x12, 0x52, 0x05, 0x7f, 0x16, 0x78, 0x17, 0x20, 0x1a, 0x3a, 0x0e, 0xf1, 0x0c, 0x46, 0x2f, 0x0f, 0x22, 0x73, 0x35, 0x10, 0x26, 0x52, 0x27, 0x0c, 0x29, 0xef, 0x00})},
			"パスワードはよろしく～戦いませんか？",
		},
		{
			"example3",
			fields{403, bytes.NewReader([]byte{0x00, 0x28, 0x17, 0x4f, 0x04, 0x3f, 0x0d, 0x6d, 0x1e, 0x3b, 0x11, 0x71, 0x09, 0x12, 0x15, 0x21, 0x06, 0x44, 0x19, 0x25, 0x18, 0xc5, 0x16, 0x57, 0x0e, 0x25, 0x21, 0x76, 0x2f, 0xcc, 0x25, 0x19, 0x36, 0x08, 0x29, 0x4a, 0x24, 0x0e, 0x2d, 0x60, 0x3e, 0x04, 0x00})},
			"そろそろ落ちます参加しま～すありがとう",
		},
		{
			"example4",
			fields{404, bytes.NewReader([]byte{0x00, 0x24, 0x15, 0x35, 0x10, 0x3a, 0x07, 0x75, 0x1f, 0x37, 0x06, 0x7f, 0x1d, 0x71, 0x01, 0xf1, 0x07, 0x52, 0x18, 0x24, 0x15, 0x2e, 0x1f, 0xf5, 0x18, 0x3b, 0x2b, 0x07, 0x22, 0x47, 0x24, 0x45, 0x37, 0x7b, 0x2b, 0xc1, 0x25, 0x14, 0x00})},
			"了解お疲れ様でした～部屋作りま～す",
		},
		{
			"Issue 22",
			fields{266, bytes.NewReader([]byte{0x00, 0x08, 0x02, 0xb2, 0x8f, 0x6b, 0x82, 0x63, 0x90, 0x5c, 0x00})},
			"ＧＭⅡ",
		},
		{
			"face",
			fields{405, bytes.NewReader([]byte{0x00, 0x08, 0x02, 0x91, 0x01, 0xcd, 0x0c, 0xd1, 0x03, 0xf4, 0x00})},
			"＾ゞ）",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MessageBodyReader{
				seq: tt.fields.seq,
				r:   tt.fields.r,
			}
			if got := m.ReadEncryptedString(); got != tt.want {
				t.Errorf("MessageBodyReader.ReadEncryptedString() = %v, want %v", got, tt.want)
			}
		})
	}
}