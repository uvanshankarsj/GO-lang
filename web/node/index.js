var express=require ("express")
var app=express()

app.get('/',(req,res)=>{
    console.log(" working")
    res.send("hi")
})

app.post('/post',(req,res)=>{
    let myjson=req.body;
     console.log(myjson)
    res.status(200).send(myjson);
})

app.listen(8009)
