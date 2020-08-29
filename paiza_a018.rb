@cameras = []

def onArea(camX, camY, mX, mY, startAngle, endAngle, rad)
    dx = mX - camX
    dy = mY - camY
    sx = Math.cos(startAngle * Math::PI / 180)
    sy = Math.sin(startAngle * Math::PI / 180)
    ex = Math.cos(endAngle * Math::PI / 180)
    ey = Math.sin(endAngle * Math::PI / 180)

    # on circle
    if dx ** 2 + dy ** 2 > rad ** 2
        return false
    end

    # when deg over 180
    if sx * ey - ex * sy > 0
        if sx * dy - dx * sy < 0 # left
            return false
        elsif ex * dy - dx * ey > 0 # right
            return false
        end
        return true
    else
        if sx * dy - dx * sy >= 0 # left
            return true
        elsif ex * dy - dx * ey <= 0 # right
            return true
        end
        return false
    end
end

def cameraAdd(x, y, t, d, r)
    @cameras << {x: x, y: y, start: t-(d/2), end: t+(d/2), rad: r}
end

def check(x, y)
    @cameras.each do |c|
        return true if onArea(c[:x], c[:y], x, y, c[:start], c[:end], c[:rad])
    end
    return false
end

first = true
camMax = 0
camCnt = 0

loop do
    line = gets
    break if line.nil?
    if first
        _, _, camMax, artMax = line.split(' ').map(&:to_i)
        first = false
    elsif camCnt < camMax
        x ,y, t, d, r = line.split(' ').map(&:to_i)
        cameraAdd(x, y, t, d, r)
        camCnt+=1
    else
        x ,y = line.split(' ').map(&:to_i)
        if check(x, y)
            puts "yes"
        else
            puts "no"
        end
    end
end