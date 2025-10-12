import express = require('express');
const router = express.Router();

router.get('/', (req: express.Request, res: express.Response) => {
    res.render('index');
});

router.get('/test', (req: express.Request, res: express.Response) => {
    res.json({
        code: 200,
        method: 'get',
        message: 'success'
    })
});

router.post('/test', (req: express.Request, res: express.Response) => {
    res.json({
        code: 201,
        method: 'post',
        message: 'created'
    })
});

router.put('/test', (req: express.Request, res: express.Response) => {
    res.json({
        code: 204,
        method: 'put',
        message: 'updated'
    })
});

router.delete('/test', (req: express.Request, res: express.Response) => {
    res.json({
        code: 200,
        method: 'delete',
        message: 'deleted'
    })
});

export default router;