def solution(a)
    map = {
        ")" => "(",
        "}" => "{",
        "]" => "[",
    }
    stack = []

    a.each_char do |c|
        case c
        when "(", "{", "["
            stack.push c
        when ")", "}", "]"
            if map[c] == stack.last
                stack.pop
            else
                return 0
            end
        end
    end

    if stack.size ==0
        return 1
    end
    0
end

p solution("{[()()]}")
p solution("([)()])")