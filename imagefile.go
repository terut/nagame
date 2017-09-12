package main

import "github.com/therecipe/qt/core"

const (
	ImageOriginal = int(core.Qt__UserRole) + 1<<iota
	ImageThumbnail
)

type ImageModel struct {
	core.QAbstractListModel

	_ func() `constructor:"init"`

	_ map[int]*core.QByteArray `property:"roles"`
	_ []*ImageFile             `property:"imageFile"`
}

type ImageFile struct {
	core.QObject

	_ string `property:"original"`
	_ string `property:"thumbnail"`
}

//func init() {
//	ImageFile_QRegisterMetaType()
//}

func (m *ImageModel) init() {
	m.SetRoles(map[int]*core.QByteArray{
		ImageOriginal:  core.NewQByteArray2("original", len("original")),
		ImageThumbnail: core.NewQByteArray2("thumbnail", len("thumbnail")),
	})

	m.ConnectData(m.data)
	m.ConnectRowCount(m.rowCount)
	m.ConnectColumnCount(m.columnCount)
	m.ConnectRoleNames(m.roleNames)
}

func (m *ImageModel) data(index *core.QModelIndex, role int) *core.QVariant {
	if !index.IsValid() {
		return core.NewQVariant()
	}

	if index.Row() >= len(m.ImageFile()) {
		return core.NewQVariant()
	}

	var img = m.ImageFile()[index.Row()]

	switch role {
	case ImageOriginal:
		{
			return core.NewQVariant14(img.Original())
		}

	case ImageThumbnail:
		{
			return core.NewQVariant14(img.Thumbnail())
		}

	default:
		{
			return core.NewQVariant()
		}
	}
}

func (m *ImageModel) rowCount(parent *core.QModelIndex) int {
	return len(m.ImageFile())
}

func (m *ImageModel) columnCount(parent *core.QModelIndex) int {
	return 1
}

func (m *ImageModel) roleNames() map[int]*core.QByteArray {
	return m.Roles()
}
