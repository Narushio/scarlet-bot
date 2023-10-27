class SnakeToCamelConverter {
    static convertKeys(obj) {
        if (typeof obj !== 'object') {
            return obj;
        }

        if (Array.isArray(obj)) {
            return obj.map(item => SnakeToCamelConverter.convertKeys(item));
        }

        const camelObj = {};
        for (const key in obj) {
            if (Object.prototype.hasOwnProperty.call(obj, key)) {
                const camelKey = SnakeToCamelConverter.snakeToCamel(key);
                camelObj[camelKey] = SnakeToCamelConverter.convertKeys(obj[key]);
            }
        }
        return camelObj;
    }

    static snakeToCamel(str) {
        return str.replace(/_([a-z])/g, (match, letter) => letter.toUpperCase());
    }
}