
local World={}
local nPredators=1
local spawnRate=500
local deathRate=50
local speedReduction=0.048
local size=8
local speed=0.2
local nPrey=130
local preySight=0.1
local shouldRenderPallette=true

local CollisionBody={}
function CollisionBody:Create(x, y, w, h)
    local this={}

    function this:isWithinScreen()
        return x+w < SCREENW() and x > 0 and y+h < SCREENH() and y > 0
    end

    function this:isAxisAlignedCollision(cb)
        return x < cb.x+cb.w and x+ w > cb.x and  y < cb.y+cb.h and h+y > cb.y
    end
    return this
end

local Behavior={ predicate=false, action = nil }
function Behavior:Create(predicate, action)
    function Behavior:eval()
        if not predicate then
            return
        end
        action()
    end
    return self
end

local Prey={}
function Prey:new(x, y)
    local this = {x=x, y=y, r=RND(size), color=10, speed=speed, sight=preySight};
    setmetatable(this, self)
    self.__index = self

    this.dir = HEADING(
        this.x,
        this.y,
        RND(SCREENW())-this.r*2,
        RND(SCREENH())-this.r*2
    )
    this.hasSeenSomethingRecently = false

    this.born=NOW()

    this.behaviors={}

    function this:render()
        CIR(this.x, this.y, this.r, this.color)
        -- RECT(this.x, this.y, this.r*2, this.r*2, this.color)
    end
    
    function this:isValidHeading()
        if this.dir ~= this.dir then
            return false
        end

        return true
    end

    function this:moveToward(hx, hy)
        local dir = HEADING(hx, hy, this.x, this.y)
        if not this:isValidHeading() then
            return
        end

        this.dir = dir
    end

    function this:move()
        if not this.isValidHeading() then
            return
        end

        if NOW()%12 == 0 then
            this.moveToward(SCREENW()/2, SCREENH()/2)
        end

        this.x = this.x+(COS(this.dir) * this.speed)
        this.y = this.y+(SIN(this.dir) * this.speed)
    end

    function this:shouldKill()
        return false
    end

    function this:spawn()
        -- if this.tick > spawnRate then
        --     this.tick = 0
        --     table.insert(World, Prey:new(this.x, this.y))
        -- end
    end

    function this:isHostile()
        return false
    end

    function this:update()
        if not CollisionBody:Create(this.x, this.y, this.r*2, this.r*2):isWithinScreen() then
            this.moveToward(RND(SCREENW()/2-this.r*2), RND(SCREENH()/2-this.r*2))
        end

        if not CollisionBody:Create(this.x, this.y, this.r*4, this.r*4):isWithinScreen() then
            this.hasSeenSomethingRecently = true
        end

        if (NOW()-this.born) % 5 == 0 and not this.hasSeenSomethingRecently then
        end

        if NOW()-this.born % 8 == 0 then
            this.hasSeenSomethingRecently = false
        end

        this:move()
        this:spawn()
    end

    return this
end

local function makePrey(n)
    for _=0, n do
        table.insert(World, Prey:new(RND(SCREENW()), RND(SCREENH())))
    end
end

function UPDATE()
    for i, c in pairs(World) do
        c.update()
        if c.shouldKill() then
            table.remove(World, i)
        end
    end
end

function RENDER()
    CLS()
    for _, c in pairs(World) do
        c.render()
    end

    if shouldRenderPallette then
        PALLETTE()
    end
end

function INIT()
    makePrey(nPrey)
end
