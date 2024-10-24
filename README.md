![Graphe](https://raw.githubusercontent.com/gabrielaravena32/graphe-app/main/build/banner.png)
<br />

# What is Graphe?

> Graphe is a **_fast and minimalistic Bible study application_**, with a focus on using the original language texts (the Hebrew Old Testament, the Septuagint and the Greek New Testament) in conjunction with an English translation. It is currently under active development.

Features include:

- Support for ESV, GNT and LXX
- Instant details while hovering original language
- Minimalistic design with built-in themes and fonts

<br/>

# Development & Building

Live development mode: `wails dev`

Build distributable app: `wails build`

Test go code: `go test ./... -v`

Currently requires a large database file that is currently not packaged with the code. However there are plans to make this downloadable.

<br/>

# Roadmap

**Verse 0.4 - Quality of life update**

- [x] Refactor Go code to include asserts
- [ ] Wails v.3
  - [x] Basic update
  - [ ] Reinstate menu & shortcuts
- [x] Remove JS Scripture functions
- [x] Replace Virtualiser with just loading 10 chapters either side
- [ ] Right click functionality
- [ ] Animations
  - [ ] Entering settings window (look at what Discord does)
- [ ] Sound effects
- [ ] Auto-download database files on start-up
  - [ ] Functionality
  - [ ] UI
- [ ] More settings updates
  - [ ] General
    - [ ] Startup - what to show
    - [ ] Logs - open, purge
  - [ ] Version / updates
    - [ ] Show version number
    - [ ] Check for updates at login
    - [ ] Application version
    - [ ] Database version
    - [ ] User login details
  - [ ] Instant details
    - [ ] UI - drag and drop available fields
    - [ ] Functionality
    - [ ] Reset
  - [ ] Shortcuts reset (see bugs below)

<br/>

**Version 0.5 – Text details update**

- [ ] LXX Capitalisation
- [ ] Parallel texts
- [ ] Bible headings
- [ ] Footnoted variant spelling/meaning
- [ ] Conjoin word highlighting
- [ ] Lexicon data
- [ ] Morph code expansions
- [ ] Cross References
- [ ] Analysis window in sidebar: top occurances in chapter (filter out regular occurances)
- [ ] Song of Songs: "SHE", "HE", "OTHERS"
- [ ] Psalm book headings
- [ ] Divine name

<br/>

**Version 0.6 – Formatting Functions update**

- [ ] Functionality, e.g...
  - [ ] N occurances
  - [ ] Overlap with another chapter
  - [ ] Morph data
  - [ ] Specific word/root (incl. original languages)
- [ ] Combination of functions
- [ ] Modal
- [ ] Saved functions
- [ ] Settings (e.g. default colours)

<br/>

**Version 0.7 – Search update**

- [ ] Functionality, e.g...
  - [ ] Range
  - [ ] Items (e.g. topic, characters, root, strongs)
  - [ ] Commands (e.g. AND, OR, GROUPING, NOT, JOIN)
- [ ] Modal
- [ ] Settings (e.g. ranges, default text)

<br/>

**Version 0.8 - Hebrew update**

- [ ] RTL text support
- [ ] Tyndale HOT
  - [ ] Data
  - [ ] Displaying sections
  - [ ] Instant details

<br/><br/>

**Improvements**

- Zoom that works just for bible window
- Setting to reset zoom to value on startup
- Scroll to specific verse (that's not at start of paragraph)
- Add more fonts
- Add more themes
- "< Close" button in sidebar for settings at small screen sizes

**Bugs**

- Scroll too fast breaks text virtualiser
- ESV psalm 119:1
- Creating shortcut with ' key (as this screws with mysql string)
- LXX go to Susanna 1 fails (Susanna 1:1 does not exist) -> use go function CreateFirstValidRef()
- Resetting shortcuts clears the display, even though it is doing correct update in db
