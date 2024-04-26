![Graphe](https://raw.githubusercontent.com/gabrielaravena32/graphe-app/main/build/banner.png)
<br />

# What is Graphe?

> Graphe is a **_fast and minimalistic Bible study application_**, with a focus on using the original language texts (the Hebrew Old Testament, the Septuagint and the Greek New Testament) in conjunction with an English translation. It is currently under active development.

Features include:

- Support for GNT and LXX
- Instant details while hovering original language
- Minimalistic design with built-in themes

<br/>

# Development & Building

Live development mode: `wails dev`

Build distributable app: `wails build`

Currently requires a large database file that is currently not packaged with the code. However there are plans to make this downloadable.

<br/>

# Roadmap

**Version 0.1 – Greek and English update**

- [x] Keyboard shortcuts
- [x] Format display of texts
- [x] Get text/passage selection working
- [x] Instant details (word hover)
- [x] GNT version
- [x] LXX version
- [x] Display app version in bottom right
- [ ] Bug fixes

<br/>

**Version 0.2 – Data update**

- [ ] Choose passage search field
- [ ] ESV version
- [ ] HOT version
- [ ] LXX capitalisation
- [ ] Auto-download database files on start-up

<br/>

**Version 0.3 – Settings update**

- [ ] Menubar
- [ ] Settings window
  - [ ] General (startup)
  - [ ] Appearence (theme, font/size)
  - [ ] Shortcuts
  - [ ] Formatting (e.g. default colours)
  - [ ] Search (e.g. ranges, default text)
  - [ ] Instant details (what is display and order?)
  - [ ] Version/updates
- [ ] Right click functionality
- [ ] Appearance modal

<br/>

**Version 0.4 – Text details update**

- [ ] Parallel texts
- [ ] Footnoted variant spelling/meaning
- [ ] Conjoin word highlighting
- [ ] Lexicon data
- [ ] Morph code expansions

<br/>

**Version 0.5 – Functions update**

- [ ] Functionality, e.g...
  - [ ] N occurances
  - [ ] Overlap with another chapter
  - [ ] Morph data
  - [ ] Specific word/root (incl. original languages)
- [ ] Combination of functions
- [ ] Modal
- [ ] Saved functions

<br/>

**Version 0.6 – Search update**

- [ ] Functionality, e.g...
  - [ ] Range
  - [ ] Items (e.g. topic, characters, root, strongs)
  - [ ] Commands (e.g. AND, OR, GROUPING, NOT, JOIN)
- [ ] Modal

<br/>

**Later Versions**

- [ ] Multiple app states (sidebar keeps them all) - ⌘1, ⌘2, ⌘3 to cycle through screens

<br/><br/>

**Random todos**

- [ ] Split `GetScriptureSections` request into multiple batches for larger ranges that perform concurrently
- [ ] Validate bible ranges passed to go functions

**Bugs**

- [ ] Version book orders not preserved for displaying texts
- [ ] Does not currently handle missing verses well (e.g. goto)
- [ ] Should hide chapters with no verses (e.g. LXX Proverbs 25-29)
- [ ] Verses split over two paragraph breaks show verse number twice

- [ ] While virtualiser is in locked mode scrollbar breaks functionality temporarily if you scroll too far away
- [ ] Resizing window does not keep verse at top (currently: only keeps block + scroll offset)
