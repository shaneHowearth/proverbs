-- +goose Up
INSERT INTO placenames (maori_name, translation, explanation) VALUES('Eaheinomauwe', E'Captain Cook\'s spelling of the Māori name for the North Island', 'Probably stands for He Ahi nō Māui (The fire of Māui), referring to the volcanoes of the central plateau');
INSERT INTO placenames (maori_name, translation, explanation) VALUES('Ekemānuka', E'Eke (to make one\'s way through); mānuka (a native tree).', 'A party of Māori who were out hunting saw a large enemy force approaching. They made their escape unseen by crouching down and creeping through the short mānuka scrub');
INSERT INTO placenames (maori_name, translation, explanation) VALUES('Eketāhuna', 'Eke (to run aground); tāhuna (shoal or sandbank).', 'This place was as far as the Makakahi River could be navigated by waka on account of the shoals.');
INSERT INTO placenames (maori_name, translation, explanation) VALUES('Epuni', 'Correctly Te Puni.', 'Te Puni was one of the leaders os the Te Āti Awa settlement of Te Whanganui-a-Tara (Wellington) after migrating south with Te Rauparaha.');
INSERT INTO placenames (maori_name, translation, explanation) VALUES('Erua', 'Two', 'The e is a particle used before digits one to nine in enumeration. Possibly a misspelling of He rua, a cave.');

-- +goose Down
