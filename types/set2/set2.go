package mcd

import (
	"github.com/kidoman/nuggets/types"
)

func init() {
	types.Register("http://upload.wikimedia.org/wikipedia/commons/f/fa/Chicken_McNuggets.jpg")
	types.Register("http://upload.wikimedia.org/wikipedia/commons/9/92/Chicken_nuggets_-_10pc.jpg")
	types.Register("http://www.wikihow.com/images/1/1c/Make-Chicken-Nuggets-Step-7.jpg")
	types.Register("does.not.exist")
}
