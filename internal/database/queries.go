package database

import "github.com/bvinc/go-sqlite-lite/sqlite3"

type GrapheQueries struct {
	gntSection            *sqlite3.Stmt
	gntWordText           *sqlite3.Stmt
	gntWordBasicInfo      *sqlite3.Stmt
	gntWordDictionary     *sqlite3.Stmt
	gntWordInflectedCount *sqlite3.Stmt

	lxxSection            *sqlite3.Stmt
	lxxWordText           *sqlite3.Stmt
	lxxWordBasicInfo      *sqlite3.Stmt
	lxxWordInflectedCount *sqlite3.Stmt

	esvSection  *sqlite3.Stmt
	esvWordText *sqlite3.Stmt
}

func prepareQueries(g *GrapheDB, db *GrapheDBConn) {
	var err error
	db.queries = GrapheQueries{}

	db.queries.gntSection, err = db.conn.Prepare(`
        SELECT ref, word_num, text, pre, post
        FROM gnt_text
        WHERE ref >= ? AND ref <= ?;
    `)
	g.check(err)

	db.queries.gntWordText, err = db.conn.Prepare(`
        SELECT text
        FROM gnt_text
        WHERE
            ref = ?
            AND word_num = ?
        LIMIT 1;
    `)
	g.check(err)

	db.queries.gntWordBasicInfo, err = db.conn.Prepare(`
        SELECT translit, english
        FROM gnt_text_info
        WHERE
            ref = ?
            AND word_num = ?
        LIMIT 1;
    `)
	g.check(err)

	db.queries.gntWordDictionary, err = db.conn.Prepare(`
		SELECT
        	t1.form, t1.gloss, t1.strong, t1.grammar,
			(
				SELECT count(*)
				FROM gnt_text_dictionary AS t2
				WHERE t2.form = t1.form
			) as count
        FROM gnt_text_dictionary AS t1
        WHERE
           	t1.ref = ?
           	AND t1.word_num = ?;
    `)
	g.check(err)

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
	g.check(err)

	// LXX
	db.queries.lxxSection, err = db.conn.Prepare(`
        SELECT ref, word_num, text, pre, post
        FROM lxx_text
        WHERE ref >= ? AND ref <= ?;
    `)
	g.check(err)

	db.queries.lxxWordText, err = db.conn.Prepare(`
        SELECT text
        FROM lxx_text
        WHERE
            ref = ?
            AND word_num = ?
        LIMIT 1;
    `)
	g.check(err)

	db.queries.lxxWordBasicInfo, err = db.conn.Prepare(`
        SELECT
            t1.translit, t1.english, t1.strongs, t1.grammar,
            t1.dictionary_form, t1.dictionary_gloss,
            (
                SELECT count(*)
                FROM lxx_text_info AS t2
                WHERE t2.dictionary_form = t1.dictionary_form
            ) as count
        FROM lxx_text_info AS t1
        WHERE
            t1.ref = ?
            AND t1.word_num = ?
        LIMIT 1;
    `)
	g.check(err)

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
	g.check(err)

	// ESV
	db.queries.esvSection, err = db.conn.Prepare(`
        SELECT ref, word_num, text, pre, post, has_instant_details
        FROM esv_text
        WHERE ref >= ? AND ref <= ?;
    `)
	g.check(err)

	db.queries.esvWordText, err = db.conn.Prepare(`
        SELECT text
        FROM esv_text
        WHERE
            ref = ?
            AND word_num = ?
        LIMIT 1;
    `)
	g.check(err)
}
