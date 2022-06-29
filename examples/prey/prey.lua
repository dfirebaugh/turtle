
local World={}
local nPredators=1
local spawnRate=500
local deathRate=50
local speedReduction=0.048
local size=8
local speed=0.4
local nPrey=50
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

local Prey={}
function Prey:Create(x, y)
    local this = {x=x, y=y, r=RND(size), color=10, speed=speed, sight=preySight, tick=0};
    this.dir = HEADING(this.x, this.y, RND(SCREENW()-SCREENW()/2)+SCREENW()/2, RND(SCREENH()-SCREENH()/2)+SCREENH()/2)
    this.hasSeenSomethingRecently = false

    this.tick = RND(180)

    function this:render()
        CIR(this.x, this.y, this.r, this.color)
        -- RECT(this.x, this.y, this.r*2, this.r*2, this.color)
    end

    function this:moveToward(hx, hy)
        this.dir = HEADING(hx, hy, this.x, this.y)
    end

    function this:move()
        this.x = this.x+(COS(this.dir) * this.speed)
        this.y = this.y+(SIN(this.dir) * this.speed)
    end

    function this:shouldKill()
        return false
    end

    function this:spawn()
        if this.tick > spawnRate then
            this.tick = 0
            table.insert(World, Prey:Create(this.x, this.y))
        end
    end

    function this:isHostile()
        return false
    end

    function this:update()
        if not CollisionBody:Create(this.x, this.y, this.r*2, this.r*2):isWithinScreen() then
            this.moveToward(SCREENW()/2, SCREENH()/2)
            this.hasSeenSomethingRecently = true
        end

        if this.tick % 50 == 0 and not this.hasSeenSomethingRecently then
            this.moveToward(RND(SCREENW()), RND(SCREENH()))
        end

        if this.tick % 26 == 0 then
            this.hasSeenSomethingRecently = false
        end

        this:move()
        this:spawn()
        this.tick = this.tick + 1
    end

    return this
end

local Predator={}
function Predator:Create(x, y)
    local this = {x=x, y=y, r=RND(size)+2, color=7, speed=speed, sight=preySight, tick=0};
    this.dir = HEADING(this.x, this.y, RND(SCREENW()-SCREENW()/2)+SCREENW()/2, RND(SCREENH()-SCREENH()/2)+SCREENH()/2)
    this.hasSeenSomethingRecently = false

    this.tick = RND(180)

    function this:render()
        CIR(this.x, this.y, this.r, this.color)
        -- RECT(this.x, this.y, this.r*2, this.r*2, this.color)
    end

    function this:moveToward(hx, hy)
        this.dir = HEADING(hx, hy, this.x, this.y)
    end

    function this:move()
        this.x = this.x+(COS(this.dir) * this.speed)
        this.y = this.y+(SIN(this.dir) * this.speed)
    end

    function this:shouldKill()
        return false
    end

    function this:spawn()
        if this.tick > spawnRate then
            -- this.tick = 0
            -- table.insert(World, Predator:Create(this.x, this.y))
        end
    end

    function this:isHostile()
        return false
    end

    function this:update()
        if not CollisionBody:Create(this.x, this.y, this.r*2, this.r*2):isWithinScreen() then
            this.moveToward(SCREENW()/2, SCREENH()/2)
            this.hasSeenSomethingRecently = true
        end

        if this.tick % 50 == 0 and not this.hasSeenSomethingRecently then
            this.moveToward(RND(SCREENW()), RND(SCREENH()))
        end

        if this.tick % 26 == 0 then
            this.hasSeenSomethingRecently = false
        end

        this:move()
        this:spawn()
        this.tick = this.tick + 1
    end

    return this
end

local function makePrey(n)
    for _=0, n do
        table.insert(World, Prey:Create(RND(SCREENW()), RND(SCREENH())))
    end
end

local function makePredators(n)
    for _=0, n do
        table.insert(World, Predator:Create(RND(SCREENW()), RND(SCREENH())))
    end
end

local function renderPallette()
    local w=SCREENW()/15
    for i= 0, 14 do
        RECT(i*w, SCREENH()-w, w, w, i)
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
    for _, c in pairs(World) do
        c.render()
    end

    if shouldRenderPallette then
        renderPallette()
    end
end

function INIT()
    makePrey(nPrey)
    makePredators(nPredators)
end
