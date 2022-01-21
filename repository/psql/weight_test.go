package psql_test

// testing the happy path only - to improve upon these tests, we could consider
// using a table test
// func TestIntegration_CreateWeightEntry(t *testing.T) {
// 	// create a NewStorage instance and run migrations
// 	cfg := config.NewConfig()
// 	storage := psql.NewStorage()

// 	err := psql.RunUpMigrations(*cfg)
// 	if err != nil {
// 		t.Errorf("test setup failed for: CreateUser, with err: %v", err)
// 		return
// 	}

// 	err = psql.LoadFixtures(*cfg)
// 	if err != nil {
// 		t.Errorf("test setup failed for: CreateUser, with err: %v", err)
// 		return
// 	}

// 	// run the test
// 	t.Run("should create a new weight entry for user with ID 2", func(t *testing.T) {
// 		newWeightEntry, err := entity.NewWeight(80, 2, 3210, 3300)
// 		if err != nil {
// 			t.Errorf("failed to run CreateUser with error: %v", err)
// 			return
// 		}

// 		// to ensure consistency we could consider adding in a static date
// 		// i.e. time.Date(insert-fixed-date-here)
// 		// creationTime := time.Now()
// 		err = storage.CreateWeightEntry(*newWeightEntry, time.Now())
// 		// assert there is no err
// 		if err != nil {
// 			t.Errorf("failed to create new user with err: %v", err)
// 			return
// 		}

// 		// now lets verify that the user is actually created using a
// 		// separate connection to the DB and pure sql
// 		db, err := sql.Open("postgres", cfg.GetDatabaseConnString())
// 		if err != nil {
// 			t.Errorf("failed to connect to database with err: %v", err)
// 			return
// 		}
// 		queryResults := []entity.Weight{}
// 		rows, err := db.Query("SELECT id, created_at, updated_at, weight, user_id, bmr, daily_caloric_intake FROM weight WHERE user_id=$1", 2)
// 		if err != nil || rows.Err() != nil {
// 			t.Errorf("this was query err: %v", err)
// 			return
// 		}

// 		for rows.Next() {
// 			weight := entity.Weight{}
// 			err := rows.Scan(&weight.ID, &weight.CreatedAt,
// 				&weight.UpdatedAt, &weight.Weight, &weight.UserID,
// 				&weight.BMR, &weight.DailyCaloricIntake,
// 			)
// 			if err != nil {
// 				t.Errorf("there was an error during scan for CreateWeight: %v", err)
// 				return
// 			}

// 			queryResults = append(queryResults, weight)
// 		}

// 		if len(queryResults) < 2 {
// 			t.Error("failed 'create new weight entry for user with id 2'")
// 			return
// 		}

// 		if queryResults[1].BMR != newWeightEntry.BMR {
// 			t.Error("failed 'create new weight entry for user with id 2'")
// 			return
// 		}
// 	})

// 	// // run some clean up, i.e. clean the database so we have a clean env
// 	// // when we run the next test
// 	t.Cleanup(func() {
// 		err := psql.RunDownMigrations(*cfg)
// 		if err != nil {
// 			if errors.Is(err, migrate.ErrNoChange) {
// 				return
// 			}
// 			t.Errorf("test cleanup failed for: CreateUser, with err: %v", err)
// 		}
// 	})
// }
