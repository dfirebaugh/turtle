local color=0
local current_sprite=0
local mousex=0
local mousey=0
local Grid=OBJECT:clone()
local Sprites={}
local Canvas=Grid:clone()
local SpritePicker=Grid:clone()

local function initialize_sprites()
    for i=0, #Sprites do
        local spr=SPR(i)
        for pi=0, #spr do
            Sprites[i].elements[pi]=spr[pi+1]
        end
    end
    Canvas.element=SHALLOWCOPY(Sprites[0])
end


print("spritely loaded...")
function INIT()
    initialize_sprites()
    Canvas.elements=SHALLOWCOPY(Sprites[0].elements)
end
function UPDATE()
    mousex, mousey=MOUSE()

    Canvas:update()
    ColorPicker:update()
    SpritePicker:update()
    for i, s in pairs(Sprites) do
        s:update()
    end
end
function RENDER()
    CLR()
    RECT(0, 0, SCREENW(), SCREENH(), 12)

    Canvas:render()
    ColorPicker:render()
    for _, s in pairs(Sprites) do
        s:render()
    end
    RECT(mousex, mousey, 2, 2, color)
    RECT(mousex+1, mousey+1, 1, 1, 1)
end



function Grid:insert(n, elem)
    self.length=n
    for i=1, n do
        self.elements[i]=elem
    end
end
function Grid:isWithinBounds(x, y)
    local row_size=math.floor(math.sqrt(self.length))
    return x > self.xoffset and x < self.xoffset+(self.size*row_size) and y > self.yoffset and y < self.yoffset+(self.size*row_size)
end
function Grid:update()
    if MOUSEL() then
        if self:isWithinBounds(mousex, mousey) then
            local x=math.floor((mousex-self.xoffset)/self.size)
            local y=math.floor((mousey-self.yoffset)/self.size)

            self:handle_lclick((y*math.floor(math.sqrt(self.length)))+x)
        end
    end
    if MOUSER() then
        if self:isWithinBounds(mousex, mousey) then
            local x=math.floor((mousex-self.xoffset)/self.size)
            local y=math.floor((mousey-self.yoffset)/self.size)

            self:handle_rclick((y*math.floor(math.sqrt(self.length)))+x)
        end
    end
end
function Grid:render()
    local i=0
    if not self then
        return
    end

    local row_size=math.floor(math.sqrt(self.length))

    -- border
    RECT(self.xoffset-1, self.yoffset-1, (row_size*self.size)+2, (row_size*self.size)+2, self.length)

    for y=0, row_size-1 do
        for x=0, row_size-1 do
            RECT((x*self.size)+self.xoffset, (y*self.size)+self.yoffset, self.size, self.size, self.elements[i])
            i=i+1
        end
    end
end

local function toNibble(n)
    if n == nil then
        return "0"
    end
    if n == 0 then
        return "0"
    end
    if n < 10 then
        return tostring(n)
    end
    if n == 10 then 
        return "a"
    end
    if n == 11 then 
        return "b"
    end
    if n == 12 then
        return "c"
    end
    if n == 13 then
        return "d"
    end
    if n == 14 then
        return "e"
    end
    if n == 15 then
        return "f"
    end
end

SpritePicker.length=36
SpritePicker.xoffset=2
SpritePicker.yoffset=94
SpritePicker.size=9.7
SpritePicker.elements={1,2,3,4,5,6,7,8,9}
function SpritePicker:handle_lclick(index)
    if index > 6 then
        return
    end
    SPRITEINDEX(index)
end
function SpritePicker:handle_rclick(index)
    if index > 6 then
        return
    end
    SPRITEINDEX(index)
end

for n=0, 3 do
    local sprite=Grid:clone()
    sprite.length=64
    sprite.xoffset=(n*10)+2
    sprite.index=n
    sprite.yoffset=95
    sprite.size=1
    sprite.elements={}
    sprite:insert(64, 0)

    function sprite:handle_lclick(index)
        current_sprite=self.index
        Canvas.elements=SHALLOWCOPY(self.elements)
    end
    function sprite:handle_rclick(index)
    end
    Sprites[n]=sprite
end

Canvas.xoffset=2
Canvas.yoffset=2
Canvas.size=10
Canvas.elements={}
Canvas:insert(64, 0)
function Canvas:out()
    local str=""
    for i=0, #self.elements-1 do
        str=str..toNibble(self.elements[i])
    end
    TOSPRITE(str)
end
function Canvas:handle_lclick(index)
    self.elements[index]=color
    Sprites[current_sprite].elements=SHALLOWCOPY(self.elements)
    self:out()
end
function Canvas:handle_rclick(index)
    color=self.elements[index]
end

ColorPicker=Grid:clone()
ColorPicker.length=16
ColorPicker.xoffset=88
ColorPicker.yoffset=8
ColorPicker.size=8
ColorPicker.elements={1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16}
function ColorPicker:handle_lclick(index)
    color=index
end
function ColorPicker:handle_rclick(index)
    color=index
end
