package formatter

const (
	invalidLine = "blabla bla"
	phpLine     = "[26-Nov-2024 14:20:15 UTC] PHP Warning:  Division by zero in /var/www/html/index.php on line 10"
	apacheLine  = "127.0.0.1 - frank [10/Oct/2024:13:55:36 -0700] \"GET /apache_pb.gif HTTP/1.0\" 200 2326"
	laravelLine = "[2024-11-26 14:20:15] local.ERROR: Something went wrong {\"user_id\":1,\"exception\":{\"key\": 1545, \"value\": 78787}}"
	yii1Line    = "2024/11/26 14:20:15 [error] [application] Something went wrong in the application."
	yii2Line    = "2024-11-26 14:20:15 [error][application] Something went wrong."
	genericLine = "block1 block2 {\"item1\": 1, \"item2\": 2}"
)
