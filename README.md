# golang-with-mongodb

find มี 5 method

    Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (cur *mongo.Cursor, err error)
    FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult
    FindOneAndDelete(ctx context.Context, filter interface{}, opts ...*options.FindOneAndDeleteOptions) *mongo.SingleResult
    FindOneAndReplace(ctx context.Context, filter interface{}, replacement interface{}, opts ...*options.FindOneAndReplaceOptions) *mongo.SingleResult
    FindOneAndUpdate(ctx context.Context, filter interface{}, update interface{}, opts ...*options.FindOneAndUpdateOptions) *mongo.SingleResult

insert มี 2 method

    InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
    InsertMany(ctx context.Context, documents []interface{}, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error)

update มี 3 method

    UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
    UpdateMany(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
    UpdateByID(ctx context.Context, id interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)

delete มี 2 method

    DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error)
    DeleteMany(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error)

---

Data Types BSON 

https://www.mongodb.com/docs/drivers/go/current/fundamentals/bson

bson.D{} ตัวอย่างใน doc เค้าใช้ , คั่นระหว่าง Key-Value 
    
    bson.D{{"foo", "bar"}, {"hello", "world"}, {"pi", 3.14159}}
    func (d primitive.D) Map() primitive.M

โค้ดตอนนำไปใช้งาน filter เพื่อ update ข้อมูล

    _id, _ := primitive.ObjectIDFromHex("6637f5e2c8849ab3cd545a21")
	filter := bson.D{{"_id", _id}}
	update := bson.D{{"$set", bson.D{{"completed", true}}}}

เปลี่ยน format ในรูป Key-Value ทำให้อ่านดูง่ายขึ้น (นิดนุง...)

    _id, _ := primitive.ObjectIDFromHex("6637f5e2c8849ab3cd545a21")
    filter := bson.D{{Key: "_id", Value: _id}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "completed", Value: true}}}}


ตัวอย่างการเริ่มเขียน golang กับ mongo: https://www.mongodb.com/docs/drivers/go/current/quick-start/

โค้ดตัวอย่าง หลักๆจะอยู่ใน code-snippets (จะเจอโค้ดส่วนอื่นๆอีก update, delete etc..): 
https://github.com/mongodb/docs-golang/blob/master/source/includes/usage-examples/code-snippets/find.go

---

MIT License

Copyright (c) 2024 Eak Netpanya

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
