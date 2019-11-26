package series

import (
	"github.com/KyleBanks/goodreads/responses/work"
)

type SeriesWork struct {
	ID           int       `xml:"id"`
	UserPosition string    `xml:"user_position"`
	Work 		 work.Work `xml:"work"`
}

type Series struct {
	ID               int     	  `xml:"id"`
	Title			 string  	  `xml:"title"`
	Description		 string	 	  `xml:"description"`
	Note			 string	 	  `xml:"note"`
	SeriesWorksCount int	 	  `xml:"series_works_count"`
	PrimaryWorkCount int	 	  `xml:"primary_work_count"`
	Numbered 		 bool	 	  `xml:"numbered"`
	SeriesWorks 	 []SeriesWork `xml:"series_works>series_work"`
}
