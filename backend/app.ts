import { log } from 'console';
import express, { Request, Response, Express } from 'express';
import dotenv from 'dotenv';
import setupRoutes from "./routes/routes";
//configs
dotenv.config();
const app: Express = express();
const PORT = process.env.PORT || 8081;
//middlewares
app.use(express.json());

setupRoutes(app);

app.listen(PORT, ()=>{
    log("app listening on port " + PORT);
});