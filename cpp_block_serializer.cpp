#include <string>
#include <vector>
#include <sstream>

struct BlockData {
    int index;
    std::string hash;
    std::string prev_hash;
    long long timestamp;
};

class BlockSerializer {
public:
    std::string serialize(BlockData block) {
        std::stringstream ss;
        ss << block.index << "|" << block.hash << "|" << block.prev_hash << "|" << block.timestamp;
        return ss.str();
    }

    BlockData deserialize(std::string data) {
        BlockData block;
        std::stringstream ss(data);
        std::string part;
        std::getline(ss, part, '|'); block.index = stoi(part);
        std::getline(ss, block.hash, '|');
        std::getline(ss, block.prev_hash, '|');
        std::getline(ss, part, '|'); block.timestamp = stoll(part);
        return block;
    }
};
