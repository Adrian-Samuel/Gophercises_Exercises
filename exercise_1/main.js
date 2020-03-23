const {readFile} = require('fs').promises;

const flagOptions = (arguments) => {
   return arguments.slice(2).reduce((options, currentOption) => {
        const [option, value] = currentOption.split("=")
        const cleanedOption = option.replace(/-/g, "")
        options[cleanedOption] = value;
        return options;
    }, {})

}

const getAnswers = async () => {

    const {filename} = flagOptions(process.argv)
  
    const getFile = await readFile(filename, "utf8");

    const rows = getFile.split(/\r?\n/g);
    const results = {
        right: 0,
        wrong: 0
    };

    for (const record of rows) {
        const [operation, answer] = record.split(',');
        const [firstNumber, secondNumber] = operation.split('+');

        const checkOperation = Number(firstNumber) + Number(secondNumber) === Number(answer);

        if (checkOperation) results["right"] += 1;
        if (!checkOperation) results["wrong"] += 1

    }

    console.log(results)
};



getAnswers();
