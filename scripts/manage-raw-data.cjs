var fs = require('fs');
var data = require('../data/wikidot/raw.json');

const orphanKeys = Object.keys(data)
  .filter((key) => {
    const current = data[key];
    switch (current.Text[current.Text.length - 1]) {
      case 'everyday':
      case 'important':
      case 'dqvc':
      case 'head':
      case 'torso':
      case 'shield':
      case 'arms':
      case 'legs':
      case 'feet':
      case 'axes':
      case 'boomerangs':
      case 'bows':
      case 'claws':
      case 'fans':
      case 'hammers':
      case 'knives':
      case 'spears':
      case 'staves':
      case 'swords':
      case 'whips':
      case 'wands':
      case 'whips':
      case 'accessories':
      case 'monster':
      case 'monster-family':
      case 'combat-quest':
      case 'story-quest':
      case 'item-quest':
      case 'misc-quest':
      case 'orphan-quest':
      case 'spell':
      case 'ability':
      case 'attribute':
      case 'skill-tree':
      case 'vocation':
      case 'misc':
      case 'town':
      case 'location':
        return false;
      default:
        return true;
    }
  })
  .forEach((key) => {
    console.log(key);
  });

// console.log(orphanKeys, orphanKeys.length);

// fs.writeFile(
//   './data/wikidot/raw.json',
//   JSON.stringify(data, null, 2),
//   'utf8',
//   (err) => console.error(err)
// );
