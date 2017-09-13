import QtQuick 2.8
import QtQuick.Controls 2.1
import QtQuick.Layouts 1.3

ApplicationWindow {
    id: window
    visible: true
    title: "Hello World Example"
    minimumWidth: 1000
    minimumHeight: 700
    color: "#111"

    RowLayout {
        id: container
        anchors.fill: parent
        spacing: 0

        function showImage(imgs) {
            leftScreen.source = imgs[1] || ""
            rightScreen.source = imgs[0]
        }

        focus: true
        Keys.onPressed: {
            if (event.key === Qt.Key_T) {
                thumbnails.visible = !thumbnails.visible
            }
        }
        Keys.onRightPressed: {
            if(presenter.hasPrev()) {
                var imgs = presenter.prevImages()
                showImage(imgs)
                thumbnails.currentIndex = presenter.getPage() - 1
            }
        }
        Keys.onLeftPressed: {
            if(presenter.hasNext()) {
                var imgs = presenter.nextImages()
                showImage(imgs)
                thumbnails.currentIndex = presenter.getPage() - 1
            }
        }

        ListView {
            id: thumbnails
            //visible: false
            spacing: 6
            width: 100
            Layout.fillHeight: true

            Component {
                id: thumbsDelegate
                Image {
                    id: thumb
                    anchors.horizontalCenter: parent.horizontalCenter
                    asynchronous: true
                    cache: false
                    source: original
                    fillMode: Image.PreserveAspectFit
                    width: 90
                    height: 100
                    Layout.preferredWidth: 100
                    Layout.preferredHeight: 100
                    MouseArea {
                        anchors.fill: parent
                        onClicked: {
                            thumbnails.currentIndex = index
                            presenter.setPage(index + 1)
                            var imgs = presenter.getImages()
                            container.showImage(imgs)
                        }
                    }
                    onStatusChanged: {
                        if (thumb.status == Image.Ready) {
                            thumb.height = thumb.implicitHeight * (parseFloat(thumb.width) / thumb.implicitWidth)
                        }
                    }
                }
            }

            model: presenter.imageModel
            delegate: thumbsDelegate
        }

        RowLayout {
            id: images
            Layout.fillWidth: true
            Layout.fillHeight: true

            DropArea {
                anchors.fill: parent
                onDropped: {
                    if (drop.hasUrls) {
                        presenter.fileDropped(drop.urls[0])
                        var imgs = presenter.getImages()
                        container.showImage(imgs)
                        thumbnails.model.modelReset()
                    }
                }
            }

            Image {
                id: leftScreen
                anchors.verticalCenter: parent.verticalCenter
                asynchronous: true
                cache: false
                source: ""
                fillMode: Image.PreserveAspectFit
                Layout.fillWidth: true
                Layout.fillHeight: true
            }
            Image {
                id: rightScreen
                anchors.verticalCenter: parent.verticalCenter
                asynchronous: true
                cache: false
                source: ""
                fillMode: Image.PreserveAspectFit
                Layout.fillWidth: true
                Layout.fillHeight: true
            }
        }
    }
}
