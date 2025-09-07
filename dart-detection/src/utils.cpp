//
// Created by hector on 22/07/25.
//
#include "../include/utils.hpp"

namespace utils {

    std::string get_env_variable(const std::string& env_variable_name) {
        const char* value = std::getenv(env_variable_name.c_str());
        if (value == nullptr) {
            return {};
        }
        return { value };
    }

}
