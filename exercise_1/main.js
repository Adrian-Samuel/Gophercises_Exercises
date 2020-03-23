const {readFile} = require('fs').promises;

const  parseFile = async () => {
    const fileName = process.argv.slice(2)[0].split("=")[1];

    const getFile = await readFile(fileName, "utf8");

    const rows = getFile.split(/\r?\n/g);
    const results = {right: 0, wrong:0};

    for(const record of rows) {
    const [operation, answer] = record.split(',');
    const [firstNumber, secondNumber] = operation.split('+');

    const mathExpression = Number(firstNumber) + Number(secondNumber) === Number(answer);

    if(mathExpression)results["right"] += 1;
    if(!mathExpression)results["wrong"] += 1

    }

    console.log(results)
};



parseFile();
