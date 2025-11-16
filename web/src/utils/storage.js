// storage.js
export const storage = {
    set(key, value) {
        localStorage.setItem(key, JSON.stringify(value));
    },
    get(key) {
        const value = localStorage.getItem(key);
        return value ? JSON.parse(value) : null;
    },
    remove(key) {
        localStorage.removeItem(key);
    }
    // 可以添加更多方法，例如清空存储等
};
