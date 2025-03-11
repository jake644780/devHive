"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const console_1 = require("console");
const express_1 = __importDefault(require("express"));
const dotenv_1 = __importDefault(require("dotenv"));
const routes_1 = __importDefault(require("./routes/routes"));
//configs
dotenv_1.default.config();
const app = (0, express_1.default)();
const PORT = process.env.PORT || 8081;
//middlewares
app.use(express_1.default.json());
(0, routes_1.default)(app);
app.listen(PORT, () => {
    (0, console_1.log)("app listening on port " + PORT);
});
