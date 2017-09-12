import QtQuick 2.8
import QtQuick.Controls 2.1
import QtQuick.Layouts 1.3

ApplicationWindow {
    id: window
    visible: true
    title: "Hello World Example"
    minimumWidth: 1000
    minimumHeight: 700

    RowLayout {
        id: container
        anchors.fill: parent
        spacing: 0

        function showImage(index) {
            leftScreen.source = presenter.getImage(index + 1).original
            rightScreen.source = presenter.getImage(index).original
        }

        focus: true
        Keys.onPressed: {
            if (event.key === Qt.Key_T) {
                thumbnails.visible = !thumbnails.visible
            }
        }
        Keys.onRightPressed: {
            thumbnails.currentIndex = thumbnails.currentIndex - 2
            showImage(thumbnails.currentIndex)
        }
        Keys.onLeftPressed: {
            thumbnails.currentIndex = thumbnails.currentIndex + 2
            showImage(thumbnails.currentIndex)
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
                            container.showImage(index)
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
                        container.showImage(0)
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
