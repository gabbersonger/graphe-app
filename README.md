![Graphe](https://raw.githubusercontent.com/gabrielaravena32/graphe-app/main/build/banner.png)
<br />

# What is Graphe?

> Graphe is a **_fast and minimalistic Bible study application_**, with a focus on using the original language texts (the Hebrew Old Testament, the Septuagint and the Greek New Testament) in conjunction with an English translation. It is currently under active development.

Features include:

- Support for ESV, GNT and LXX
- Instant details while hovering original language
- Minimalistic design with built-in themes

<br/>

# Development & Building

Live development mode: `wails dev`

Build distributable app: `wails build`

Currently requires a large database file that is currently not packaged with the code. However there are plans to make this downloadable.

<br/>

# Roadmap

**Version 0.2 – Data update**

- [x] Standardise instant details
- [ ] Choose passage
  - [ ] UI
  - [ ] Search field
  - [ ] Missing chapters (e.g. LXX Prov 25-29)
- [ ] ESV version
  - [x] Data
  - [x] Displaying sections
  - [ ] Instant details
- [ ] LXX updates
  - [ ] Capitalisation
  - [ ] Punctuation in psalms
- [ ] Auto-download database files on start-up
  - [ ] Functionality
  - [ ] UI
- [ ] Text UI updates
  - [x] Fix basic punctuation
  - [ ] Chapters starting in middle of paragraph
  - [ ] Verses split over multiple paragraphs

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

**Verse 0.4 - Quality of life update**

- [ ] Improve virtualiser speed
- [ ] Animations
- [ ] Sound effects

<br/>

**Version 0.5 – Text details update**

- [ ] Parallel texts
- [ ] Footnoted variant spelling/meaning
- [ ] Conjoin word highlighting
- [ ] Lexicon data
- [ ] Morph code expansions
- [ ] Cross References
- [ ] Analysis window in sidebar: top occurances in chapter (filter out regular occurances)

<br/>

**Version 0.6 – Functions update**

- [ ] Functionality, e.g...
  - [ ] N occurances
  - [ ] Overlap with another chapter
  - [ ] Morph data
  - [ ] Specific word/root (incl. original languages)
- [ ] Combination of functions
- [ ] Modal
- [ ] Saved functions

<br/>

**Version 0.7 – Search update**

- [ ] Functionality, e.g...
  - [ ] Range
  - [ ] Items (e.g. topic, characters, root, strongs)
  - [ ] Commands (e.g. AND, OR, GROUPING, NOT, JOIN)
- [ ] Modal

<br/>

**Version 0.8 - Hebrew update**

- [ ] RTL text support
- [ ] Tyndale HOT
  - [ ] Data
  - [ ] Displaying sections
  - [ ] Instant details

<br/>

**Later Versions**

- [ ] Multiple app states (sidebar keeps them all) - ⌘1, ⌘2, ⌘3 to cycle through screens

<br/><br/>

**Bugs**

- [ ] Does not currently handle missing verses well (e.g. goto)
- [ ] While virtualiser is in locked mode scrollbar breaks functionality temporarily if you scroll too far away
- [ ] Resizing window does not keep verse at top (currently: only keeps block + scroll offset)
- [ ] On shutdown, database is not being closed properly
