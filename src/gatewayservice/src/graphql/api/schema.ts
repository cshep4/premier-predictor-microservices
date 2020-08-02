import * as fs from "fs";
import * as path from "path";

const walk = (dir: string): string[] => {
    let results: string[] = [];
    const list = fs.readdirSync(dir);

    list.forEach((file: string) => {
        file = path.resolve(dir, file);
        const stat = fs.statSync(file);
        if (stat && stat.isDirectory()) {
            results = results.concat(walk(file));
            return;
        }

        if (path.extname(file) == '.graphql') {
            results.push(fs.readFileSync(file).toString('utf8'));
        }
    });

    return results;
};

const typeDefs: any = walk(__dirname);

export default typeDefs;
