process.stdin.setEncoding('utf8');
process.stdin.on('data', data => {
    const n = data.split(" ");
    for(let i=0; i<=n[1]-1; i++) {
        const out = '*'.repeat(n[0])
        process.stdout.write(`${out}\n`);
    }
});
