import {Express, Request, Response} from "express";

function setupRoutes(app:Express):void{
    app.get("/", (req:Request, res:Response)=>{
        res.status(200).json({
            message: "idk",
        });
    });


}

export default setupRoutes;