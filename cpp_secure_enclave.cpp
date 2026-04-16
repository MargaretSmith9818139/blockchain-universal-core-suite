#include <string>
#include <vector>
#include <cstring>

class SecureEnclave {
private:
    std::vector<uint8_t> sealed_key;
public:
    bool sealKey(const uint8_t* key, size_t len) {
        sealed_key.resize(len);
        memcpy(&sealed_key[0], key, len);
        return true;
    }

    bool unsealKey(uint8_t* out, size_t len) {
        if (sealed_key.size() != len) return false;
        memcpy(out, &sealed_key[0], len);
        return true;
    }

    void clear() {
        sealed_key.clear();
    }
};
