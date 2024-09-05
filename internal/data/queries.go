package data

import "github.com/bvinc/go-sqlite-lite/sqlite3"

type DataDBQueries struct {
	gntSection            *sqlite3.Stmt
	gntWordText           *sqlite3.Stmt
	gntWordBasicInfo      *sqlite3.Stmt
	gntWordDictionary     *sqlite3.Stmt
	gntWordInflectedCount *sqlite3.Stmt

	lxxSection            *sqlite3.Stmt
	lxxWordText           *sqlite3.Stmt
	lxxWordBasicInfo      *sqlite3.Stmt
	lxxWordInflectedCount *sqlite3.Stmt

	esvSection         *sqlite3.Stmt
	esvWordBasicInfo   *sqlite3.Stmt
	esvWordStrongsInfo *sqlite3.Stmt
}

func (d *DataDB) prepareQueries(db *DataDBConn) {
	var err error
	db.queries = DataDBQueries{}

	db.queries.gntSection, err = db.conn.Prepare(`
        SELECT ref, word_num, text, pre, post, 1 as has_instant_details
        FROM gnt_text
        WHERE ref >= ? AND ref <= ?;
    `)
	d.assert(err == nil, "Error preparing query `gntSection`")

	db.queries.gntWordText, err = db.conn.Prepare(`
        SELECT text
        FROM gnt_text
        WHERE
            ref = ?
            AND word_num = ?
        LIMIT 1;
    `)
	d.assert(err == nil, "Error preparing query `gntWordText`")

	db.queries.gntWordBasicInfo, err = db.conn.Prepare(`
        SELECT translit, english
        FROM gnt_text_info
        WHERE
            ref = ?
            AND word_num = ?
        LIMIT 1;
    `)
	d.assert(err == nil, "Error preparing query `gntWordBasicInfo`")

	db.queries.gntWordDictionary, err = db.conn.Prepare(`
		SELECT
        	t1.form, t1.gloss, t1.strong, t1.grammar,
			(
				SELECT count(*)
				FROM gnt_text_dictionary AS t2
				WHERE t2.form = t1.form
			) AS count
        FROM gnt_text_dictionary AS t1
        WHERE
           	t1.ref = ?
           	AND t1.word_num = ?;
    `)
	d.assert(err == nil, "Error preparing query `gntWordDictionary`")

	db.queries.gntWordInflectedCount, err = db.conn.Prepare(`
        SELECT count(*)
        FROM gnt_text
        WHERE text = (
        	SELECT text
         	FROM gnt_text
          	WHERE
           		ref = ?
             	AND word_num = ?
            LIMIT 1
        )
        LIMIT 1;
    `)
	d.assert(err == nil, "Error preparing query `gntWordInflectedCount`")

	// LXX
	db.queries.lxxSection, err = db.conn.Prepare(`
        SELECT ref, word_num, text, pre, post, 1 as has_instant_details
        FROM lxx_text
        WHERE ref >= ? AND ref <= ?;
    `)
	d.assert(err == nil, "Error preparing query `lxxSection`")

	db.queries.lxxWordText, err = db.conn.Prepare(`
        SELECT text
        FROM lxx_text
        WHERE
            ref = ?
            AND word_num = ?
        LIMIT 1;
    `)
	d.assert(err == nil, "Error preparing query `lxxWordText`")

	db.queries.lxxWordBasicInfo, err = db.conn.Prepare(`
        SELECT
            t1.translit, t1.english, t1.strongs, t1.grammar,
            t1.dictionary_form, t1.dictionary_gloss,
            (
                SELECT count(*)
                FROM lxx_text_info AS t2
                WHERE t2.dictionary_form = t1.dictionary_form
            ) AS count
        FROM lxx_text_info AS t1
        WHERE
            t1.ref = ?
            AND t1.word_num = ?
        LIMIT 1;
    `)
	d.assert(err == nil, "Error preparing query `lxxWordBasicInfo`")

	db.queries.lxxWordInflectedCount, err = db.conn.Prepare(`
        SELECT count(*)
        FROM lxx_text
        WHERE text = (
            SELECT text
            FROM lxx_text
            WHERE
                ref = ?
                AND word_num = ?
            LIMIT 1
        )
        LIMIT 1;
    `)
	d.assert(err == nil, "Error preparing query `lxxWordInflectedCount`")

	// ESV
	db.queries.esvSection, err = db.conn.Prepare(`
        SELECT ref, word_num, text, pre, post, has_instant_details
        FROM esv_text
        WHERE ref >= ? AND ref <= ?;
    `)
	d.assert(err == nil, "Error preparing query `esvSection`")

	db.queries.esvWordBasicInfo, err = db.conn.Prepare(`
        SELECT
            text, (
                SELECT count(*)
                FROM esv_text AS t2
                WHERE LOWER(t2.text) = LOWER(t1.text)
            ) AS count
        FROM esv_text AS t1
        WHERE
            ref = ?
            AND word_num = ?
        LIMIT 1;
    `)
	d.assert(err == nil, "Error preparing query `esvWordBasicInfo`")

	db.queries.esvWordStrongsInfo, err = db.conn.Prepare(`
        SELECT strongs, (
            SELECT count(*)
            FROM esv_text_strongs AS t2
            WHERE t2.strongs = t1.strongs
        ) AS count
        FROM esv_text_strongs AS t1
        WHERE
            ref = ?
            AND word_num = ?;
    `)
	d.assert(err == nil, "Error preparing query `esvWordStrongsInfo`")
}
