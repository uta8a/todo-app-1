const { exec } = require('node:child_process');

exports.execute = async (req, res) => {
    const out = exec('perl get.pl', (err, stdout, stderr) => {
        if (err) {
            console.error(err);
            return res.status(500
            ).send(err);
        }
        console.log(stdout);
        res.status(200).send(stdout);
    });
    if (!out.killed) {
        return;
    }
    res.status(200).send('hello world.');
}
