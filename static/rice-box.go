package static

import (
	"time"

	"github.com/GeertJohan/go.rice/embedded"
)

func init() {

	// define files
	file2 := &embedded.EmbeddedFile{
		Filename:    "invitation.tmpl",
		FileModTime: time.Unix(1551777482, 0),

		Content: string("プロラボアカウントに招待されました。\n\n下記URLにアクセスしてユーザー登録を完了してください。\n\n{{.RegistrationURL}}\n\n-----------------------------------------------\n久留米高専 プログラミングラボ部\nhttps://kurume-nct.com\n-----------------------------------------------\n"),
	}

	// define dirs
	dir1 := &embedded.EmbeddedDir{
		Filename:   "",
		DirModTime: time.Unix(1551776820, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file2, // "invitation.tmpl"

		},
	}

	// link ChildDirs
	dir1.ChildDirs = []*embedded.EmbeddedDir{}

	// register embeddedBox
	embedded.RegisterEmbeddedBox(`emails`, &embedded.EmbeddedBox{
		Name: `emails`,
		Time: time.Unix(1551776820, 0),
		Dirs: map[string]*embedded.EmbeddedDir{
			"": dir1,
		},
		Files: map[string]*embedded.EmbeddedFile{
			"invitation.tmpl": file2,
		},
	})
}
