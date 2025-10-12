"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const express = require("express");
const router = express.Router();
router.get('/', (req, res) => {
    res.render('index');
});
router.get('/test', (req, res) => {
    res.json({
        code: 200,
        method: 'get',
        message: 'success'
    });
});
router.post('/test', (req, res) => {
    res.json({
        code: 201,
        method: 'post',
        message: 'created'
    });
});
router.put('/test', (req, res) => {
    res.json({
        code: 204,
        method: 'put',
        message: 'updated'
    });
});
router.delete('/test', (req, res) => {
    res.json({
        code: 200,
        method: 'delete',
        message: 'deleted'
    });
});
exports.default = router;
//# sourceMappingURL=index.js.map