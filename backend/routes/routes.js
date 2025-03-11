"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
function setupRoutes(app) {
    app.get("/", (req, res) => {
        res.status(200).json({
            message: "idk",
        });
    });
}
exports.default = setupRoutes;
